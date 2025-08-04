# AI Interview Trainer - API

## Using Go Fiber

**DEMO:** https://www.youtube.com/watch?v=ef2ivitjiBU

_Frontend:_ https://github.com/dcrebbin/ai-interview-trainer-frontend

This api is for Up It APS made with [Go Fiber](https://docs.gofiber.io/) 
 
Up it APS is a WIP AI powered interview training platform for the Australian Public Service

However, can be used as a baseline to create an interview training platform for any industry. 

## AI Providers Supported

_Note:_ Vertex AI isn't fully supported as proper auth hasn't been integrated. However local testing can occur via a hacky workaround mentioned below

### Text Generation

- [OpenAi (GPT3.5, GPT4, etc)](https://platform.openai.com/docs/api-reference/chat)
- [Vertex (Palm, Gemini Pro etc)](https://console.cloud.google.com/vertex-ai/generative)

### Text to Speech

- [OpenAi (TTS 1, TTS 1 HD)](https://platform.openai.com/docs/api-reference/audio/createSpeech)
- [Vertex](https://console.cloud.google.com/vertex-ai/generative)
- [ElevenLabs](https://elevenlabs.io/docs/api-reference/text-to-speech)
- [Unreal Speech](https://docs.unrealspeech.com/)

### Speech to Text

- [OpenAi (Whisper 1)](https://platform.openai.com/docs/api-reference/audio/createTranscription)
- [Vertex](https://console.cloud.google.com/vertex-ai/generative)

## Setup

1. [Install GO](https://go.dev/doc/install) 

1. [Install gcloud CLI](https://cloud.google.com/sdk/docs/install)

1. Create a gcloud project and enable a bunch of things, etc etc

1. Go get

1. Create a .env using the env.example file

## Swagger

_Not fully implemented_

http://127.0.0.1:8080/swagger/index.html

## Authentication

This allows you to deploy to gcp

> gcloud auth login

Need to use auth quickly to use vertex ai?

> gcloud auth print-access-token

## Deploy

(make sure to be within the root directory ./)

> gcloud run deploy --source .
