-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id            UUID NOT NULL DEFAULT gen_random_uuid(),                              -- Уникальный идентификатор ползователя
    token         TEXT NOT NULL,                                                        -- token токен
    expires_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now() + interval '1 day',   -- Когда истекает токен
    is_revoked    BOOLEAN NOT NULL DEFAULT false,                                       -- Является ли отозванным
    user_id       UUID NOT NULL,                                                        -- Идентификатор пользователя
    created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),                      -- Когда создан
    PRIMARY KEY(id),
    UNIQUE(token)
    );

COMMENT ON TABLE refresh_tokens IS 'Таблица токенов';

COMMENT ON COLUMN refresh_tokens.id IS 'Уникальный идентификатор ползователя';
COMMENT ON COLUMN refresh_tokens.token IS 'Значение токена';
COMMENT ON COLUMN refresh_tokens.expires_at IS 'Когда истекает токен';
COMMENT ON COLUMN refresh_tokens.is_revoked IS 'Является ли отозванным';
COMMENT ON COLUMN refresh_tokens.user_id IS 'Идентификатор ползователя';
COMMENT ON COLUMN refresh_tokens.created_at IS 'Когда создан';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS refresh_tokens;
-- +goose StatementEnd