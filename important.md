# Preeweek: Prompt engineering and system design 
- Design design design. think of you project as a system and as a free language programming.
- Use project files! lile examples.md. Divide examples to sub-context examples (E.g for each state!) Include scoring for the examples, to provide adjustments for what we wish.
- Use prompt / example enhancement: use a model to create a large number of those. Eventually when I have the prompt.md file or example file(s) ready, I can ask for improvement ontop of them.
- Use state / machine design for your agent
- Include special prompts / cases to debug.

- ![Rought design sketch](RoughSketch.png).
   Here you can note that when aimming for a certain design, alot of the time you need more than the main thread of the agent providing the service, and more than the guradrails in the in-out of the system, you need a logic layer that "saves" the state of the conversasion, and understands it rather than keep sending all the conversasion again back&forth.
   Additionally you may want a second agent summarising the inputs or the responses. Maybe an evaluation layer to consider whether the output matches the criteria.
- ![Detailed top layer design](DetailedDesign.png)

# Week 0: Code writing.
## Writing
- First, keep in mind the prompt engineering from last section. This way your request and identifying the minimum viable products / processes would be clearer.
- Then keep in mind that the gen-ai agents are powerful, but they need a plan to adhere to. Create a plan from your design, *critique it*, then follow up with the coding agents, creating the rules to base their code on and create the files that follow the plan.
- Make sure to test, to "mess around" and "dirty your hands" at the process of making the codes connect, and the errors work, while sticking to your plan.

## Data
### prep.
- This is defined as the process of *transforming raw data*, into a _format_ that can be effiecient and effective for model training. Remember that even in those big project you don't want an overfit or a bias. Sticking to a uniform data is key. (Similar to the state definition)
- The input data and output data has to match the end-user needs eventually.
- Models work properly when introduced with a well defined data structure. Like xml, or json, etc. it improves pattern matching.
- raw data processing is required, in the simplest matter. cleaning the data from noisy text or incoherent input might improve results alot: eg. "URGENT!!! do this and that...!" -> "in regards to this, do this, then follow with that"
- The same is true with RAG (file intoduction). files should be ordered and coherent.

### Quality
- Accuracy: truely reflect true world conditions. (Ex. restuarant menu prices)
- Completeness : all necessary data points should be present (Ex. recipe : ingridients & how to make...)
- Consistency: does not self contridict , coherent.
- Timeliness : data is up to date

### Best practices
Cleansing:
- Remove duplicates
- handle missing
- remove noise
- handle outliers
Processing:
- Encoding categorical variables
- scaling numerical data : normalise, stardise, use units
- splitting the data into training, validation and test sets.
Feature engineering:
- Highlight for the agent the features that you need or wish to include in free language.
Data labeling & annotation:
- unlike encoding categories, can define quality level for pronunciation for example.
- can define error types. 
- can define word difficulty etc.
- In short this is like critiquing the output and labeling into categories that are of interesnt and not only factual.
Data Privacy:
- handle the user privacy.
- don't keep irrelevent or private information.
- don't share information that is private the might was saved.

Document everything. It is very good for structuring, inputing, and tracking process.

## Coding from scratch
- There will always be a payed / better tool to autocomplete your writings and have clever suggestions, but the true power of this tool is having a clear plan in mind, knowing the possibilities, and simply dictating the road ahead.

# Week 1: Developement from scratch
Alot was covered this week.
We got the API backend going, via the technical analysis.
Same with the frontend, and lovable.
-> Remember that creating the basics doesn't cut it on the long run, you have to be able to engineer the desgin to be able to hold features and updates!! And then successfully apply those features inside, can be a little different when using AI generated templates.

We covered a realy nice and strong basis for running Ollama-server.
Using docker to connnect to the ollama server, composing, then pulling and connecting with the lamma API (in the terminal).
This is an important note, after having a back and front you want an agent or at least a machine to communicate with and run locally, or at least to connect localy.
We had alot of time covered figuring out how to run the server, then pulling the ollama service and model to it, and only then communicating with it.
We had a tough go at creating a megaservice and an envalope for this ollama server, it is important to get the data-streaming-I/O objects right at everyturn, but we got the components working.

Towards the end of the week of classes, an additional implementaion of language-dictionary was (will) added to the library. This is an important lesson of recapping the tools that were covered this week, connecting with an agent, and adding major features to your design.

%% TODO: Add this language dictionary
%% TODO: (Optional) Cover the backend-tests
%% TODO: (Optional) Cover the handle_response function

%% TODO: **IMPORTANT** (SOON) Design an agent for internal use.