package ai_model

type ElevenLabsRequest struct {
	OptimizeStreamingLatency int                     `json:"optimize_streaming_latency"`
	Text                     string                  `json:"text"`
	ModelID                  string                  `json:"model_id"`
	VoiceSettings            ElevenLabsVoiceSettings `json:"voice_settings"`
}

type ElevenLabsVoiceSettings struct {
	Stability       float32 `json:"stability"`
	SimilarityBoost float32 `json:"similarity_boost"`
}
