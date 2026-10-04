package log

import "github.com/TimofeyAndriyanov/AureliaReadsBackend/domain/entities"

type Repository interface {
	Add(level entities.LogLevel, tag string)
}
