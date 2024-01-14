package ai_model

type GoogleGenerationConfig struct {
	Temperature     float32  `json:"temperature"`
	TopP            float32  `json:"topP"`
	TopK            int64    `json:"topK"`
	MaxOutputTokens int64    `json:"maxOutputTokens"`
	StopSequences   []string `json:"stopSequences"`
}

type GoogleRequest struct {
	Contents         []GoogleRequestContent `json:"contents"`
	SafetySettings   []GoogleRequestSafety  `json:"safetySettings"`
	GenerationConfig GoogleGenerationConfig `json:"generationConfig"`
}

type GoogleRequestSafety struct {
	Category  string `json:"category"`
	Threshold string `json:"threshold"`
}

type GoogleField struct {
	Content string `json:"content"`
}

type GoogleRequestContent struct {
	Parts []GoogleRequestPart `json:"parts"`
}
type GoogleRequestPart struct {
	Text string `json:"text"`
}

type GoogleResponseError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"error"`
}

type GoogleResponseCandidate struct {
	Content GoogleRequestContent `json:"content"`
}

type GoogleResponse struct {
	Candidates []GoogleResponseCandidate `json:"candidates"`
}

type GoogleVertexAiAudioRequestInput struct {
	Text string `json:"text"`
}
type GoogleVertexAiAudioRequestVoice struct {
	LanguageCode string `json:"languageCode"`
	Name         string `json:"name"`
}

type GoogleVertexAiAudioRequestAudioConfig struct {
	AudioEncoding string  `json:"audioEncoding"`
	SpeakingRate  float32 `json:"speakingRate"`
}
type GoogleVertexAiRequest struct {
	Input       GoogleVertexAiAudioRequestInput       `json:"input"`
	Voice       GoogleVertexAiAudioRequestVoice       `json:"voice"`
	AudioConfig GoogleVertexAiAudioRequestAudioConfig `json:"audioConfig"`
}

type GoogleVertexAiAudioResponse struct {
	AudioContent string `json:"audioContent"`
}

type GoogleVertexAiSpeechToTextRequest struct {
	Config GoogleVertexAiSpeechToTextRequestConfig `json:"config"`
	Audio  GoogleVertexAiSpeechToTextAudio         `json:"audio"`
}

type GoogleVertexAiSpeechToTextRequestConfig struct {
	LanguageCode          string `json:"languageCode"`
	EnableWordTimeOffsets bool   `json:"enableWordTimeOffsets"`
	EnableWordConfidence  bool   `json:"enableWordConfidence"`
	Model                 string `json:"ai_model"`
	Encoding              string `json:"encoding"`
	SampleRateHertz       int    `json:"sampleRateHertz"`
	AudioChannelCount     int    `json:"audioChannelCount"`
}

type GoogleVertexAiSpeechToTextAudio struct {
	Content string `json:"content"`
}

type GoogleWord struct {
	StartTime  string  `json:"startTime"`
	EndTime    string  `json:"endTime"`
	Word       string  `json:"word"`
	Confidence float64 `json:"confidence"`
}

type GoogleAlternative struct {
	Transcript string       `json:"transcript"`
	Confidence float64      `json:"confidence"`
	Words      []GoogleWord `json:"words"`
}

type GoogleVertexAiSpeechToTextResponseResults struct {
	Alternatives  []GoogleAlternative `json:"alternatives"`
	ResultEndTime string              `json:"resultEndTime"`
	LanguageCode  string              `json:"languageCode"`
}

type GoogleVertexAiSpeechToTextResponse struct {
	VertexAiSpeechToTextResponseResults []GoogleVertexAiSpeechToTextResponseResults `json:"results"`
	TotalBilledTime                     string                                      `json:"totalBilledTime"`
	RequestId                           string                                      `json:"requestId"`
}
