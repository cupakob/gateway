package tokenutils

import (
	"github.com/basemind-ai/monorepo/shared/go/db"
	"github.com/rs/zerolog/log"
	"github.com/tiktoken-go/tokenizer"
)

var modelEncodingMap = map[db.ModelType]tokenizer.Encoding{
	db.ModelTypeGpt35Turbo:    tokenizer.Cl100kBase,
	db.ModelTypeGpt35Turbo16k: tokenizer.Cl100kBase,
	db.ModelTypeGpt4:          tokenizer.Cl100kBase,
	db.ModelTypeGpt432k:       tokenizer.Cl100kBase,
}

func GetPromptTokenCount(prompt string, modelType db.ModelType) int32 {
	encoding := modelEncodingMap[modelType]
	enc, err := tokenizer.Get(encoding)
	if err != nil {
		log.Error().Err(err).Msg("failed to get prompt token count")
	}

	ids, _, _ := enc.Encode(prompt)
	return int32(len(ids))
}
