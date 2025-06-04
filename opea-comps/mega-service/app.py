# Taken from "https://opea-project.github.io/latest/GenAIComps/README.html"
# This file initiates a mega service from scratch.


from comps import MicroService, ServiceOrchestrator
from comps.cores.mega.constants import ServiceType, ServiceRoleType 
from comps.cores.proto.api_protocol import ChatCompletionRequest, ChatCompletionResponse
from comps.cores.proto.api_protocol import ChatMessage, UsageInfo
from fastapi import HTTPException

from starlette.responses import StreamingResponse
import json

import os

EMBEDDING_SERVICE_HOST_IP = os.getenv("EMBEDDING_SERVICE_HOST_IP", "0.0.0.0")
EMBEDDING_SERVICE_PORT = os.getenv("EMBEDDING_SERVICE_PORT", 6000)
LLM_SERVICE_HOST_IP = os.getenv("LLM_SERVICE_HOST_IP", "0.0.0.0")
LLM_SERVICE_PORT = os.getenv("LLM_SERVICE_PORT", 8008)


class ExampleService:
    def __init__(self, host="0.0.0.0", port=8000):
        print(f"A service is being created at host {host} and port {port}")
        self.host = host
        self.port = port
        self.endpoint = "/v1/example-service"
        self.megaservice = ServiceOrchestrator()
        print(f'A service was initiated on "{self.endpoint}"')

    def add_remote_service(self):
        print(f"A remote service is being attached...")
        # Note: Removed emedding for now.
        # embedding = MicroService(
        #     name="embedding",
        #     host=EMBEDDING_SERVICE_HOST_IP,
        #     port=EMBEDDING_SERVICE_PORT,
        #     endpoint="/v1/embeddings",
        #     use_remote_service=True,
        #     service_type=ServiceType.EMBEDDING,
        # )
        llm = MicroService(
            name="llm",
            host=LLM_SERVICE_HOST_IP,
            port=LLM_SERVICE_PORT,
            endpoint="/v1/chat/completions",
            use_remote_service=True,
            service_type=ServiceType.LLM,
        )
        # self.megaservice.add(embedding).add(llm)
        # self.megaservice.flow_to(embedding, llm)
        self.megaservice.add(llm)

    async def handle_request(self, request: ChatCompletionRequest) -> ChatCompletionResponse:
        try:
            # print("Incoming request:", request.model_dump())

            # Schedule the request — returns a dict of responses and the DAG
            response_map, _ = await self.megaservice.schedule(request.model_dump())

            # Extract the StreamingResponse object from the map
            raw_response = response_map.get("llm/MicroService")

            if not isinstance(raw_response, StreamingResponse):
                raise HTTPException(status_code=500, detail="Invalid response type from LLM microservice.")

            # Read and decode the streamed response
            body = b"".join([chunk async for chunk in raw_response.body_iterator])
            print("RAW BODY:", body.decode(errors="ignore"))

            if not body:
                raise HTTPException(status_code=502, detail="LLM service returned an empty response body.")

            try:
                response_json = json.loads(body)
            except json.JSONDecodeError as e:
                print("Raw body was:", body.decode(errors="ignore"))
                raise HTTPException(status_code=502, detail=f"Failed to decode LLM response: {e}")

            # Convert to ChatCompletionResponse model
            response = ChatCompletionResponse(**response_json)

            # Ensure there is at least one assistant message
            if not response.choices:
                response.choices = [{
                    "message": ChatMessage(role="assistant", content="I'm not sure how to respond to that."),
                    "finish_reason": "stop",
                    "index": 0,
                }]

            # Add default usage if none is provided
            if response.usage is None:
                response.usage = UsageInfo(
                    prompt_tokens=10,
                    completion_tokens=15,
                    total_tokens=25,
                )

            print("Final structured response:", response.model_dump())
            return response

        except Exception as e:
            print(f"Error during async handling of request: {e}")
            raise HTTPException(status_code=500, detail=f"Internal Server Error: {str(e)}")


    def start(self):
        print(f"Starting the mega service...")
        self.service = MicroService(
            self.__class__.__name__,
            service_role=ServiceRoleType.MEGASERVICE,
            host=self.host,
            port=self.port,
            endpoint=self.endpoint,
            input_datatype=ChatCompletionRequest,
            output_datatype=ChatCompletionResponse,
        )

        self.service.add_route(self.endpoint, self.handle_request, methods=["POST"])

        self.service.start()

example = ExampleService()
example.add_remote_service()
example.start()