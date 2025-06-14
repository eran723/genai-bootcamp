# OPEA COMPS
OPEA components are a set of reusable components designed to be used in OPEA appllications.
This is used to opperate AI modules locally.

## Read More
### Gen AI Bootcamp
https://docs.google.com/document/d/1KVDTDF4t8VtI69F5KMo67KoTBXgVhsd2O9hK-uPh2rA/edit?tab=t.m24gv8gcvfj3
https://github.com/ExamProCo/gen-ai-training-day-workshops
https://github.com/ExamProCo/gen-ai-training-day-workshops/blob/main/002__intermediate-workshop/opea.sh
https://github.com/ExamProCo/GenAI-Essentials/blob/main/opea/basic/opea.sh
### OPEA Components
https://github.com/opea-project/GenAIComps
https://github.com/opea-project/GenAIComps/tree/main/comps/third_parties/ollama
### OPEA Project Documentation
https://opea-project.github.io/latest/microservices/index.html#llms-microservice

## Technical Restrictions
Since we are using our local machine in this GenAI course, we might not be able to run suffisticated processors like TGI or gaudi2.
So we proceed with the locall ollama container.
We will use dockers to run everything.

## Running Ollama Third-Party Service
### Download (Pull) a model
curl http://localhost:8008/api/pull -d '{
  "model": "llama3.2:1b"
}'
curl http://localhost:8008/api/pull -d '{
  "model": "tinyllama:1.1b"
}'

### # Choosing a modules
You can get the model_id that ollama will launch from the [Ollama Library] https://ollama.com/library
e.g https://ollama.com/library/llama3.2

#### Linux
Get your IP address
'''sh
sudo apt install net-tools 
ifconfig
'''
Or try it this way : '$(hostname -I | awk '{print $1}')'

##### Configuration Run
(host found on eth0, after running ifconfig)
'''sh
cd opea-comps/
export NO_PROXY=localhost
export HOST_IP=$(hostname -I | awk '{print $1}')
export LLM_ENDPOINT_PORT=8008 
export LLM_MODEL_ID="llama3.2:1b"
docker compose up
'''

### API access
https://github.com/ollama/ollama/blob/main/docs/api.md

https://github.com/opea-project/GenAIComps/blob/main/comps/third_parties/ollama/README.md

Send an application/json request to the API endpoint of Ollama to interact.

Once the Ollama server is running we can make API calls to the ollama API

#### Generate request
##### llama3.2:1b
curl http://localhost:8008/api/generate -d '{
  "model": "llama3.2:1b",
  "prompt":"Why is the sky blue?"
}'
##### tinyllama:1.1b
curl http://localhost:8008/api/generate -d '{
  "model": "tinyllama:1.1b",
  "prompt":"Hello! Which model are you?"
}'

Note that you should explicitly mention 'ports:' in the container description. having localport <-> internal docker port.

## OPEA Examples
Examples on running OPEA Micro/Mega Services.