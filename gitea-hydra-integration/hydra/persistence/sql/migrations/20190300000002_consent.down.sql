ALTER TABLE hydra_oauth2_consent_request DROP COLUMN forced_subject_identifier;
ALTER TABLE hydra_oauth2_authentication_request_handled DROP COLUMN forced_subject_identifier;
DROP TABLE hydra_oauth2_obfuscated_authentication_session;
