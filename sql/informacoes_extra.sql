-- Table: marcelo_result_final.informacoes_extra

-- DROP TABLE marcelo_result_final.informacoes_extra;

CREATE TABLE marcelo_result_final.informacoes_extra
(
  id text,
  codigo text,
  nome text,
  material text,
  sinonimo text,
  volume text,
  metodo text,
  volumelab text,
  volumeminimo text,
  estabiliadeambienteqtd text,
  estabilidadeambientetipo text,
  estabilidaderefrigeradaqtd text,
  estabilidaderefrigeradatipo text,
  estabilidadefreezerqtd text,
  estabilidadefreezertipo text,
  rotina text,
  resultado text,
  temperatura text,
  coleta text,
  codigosus text,
  cbhpm text,
  esoterico text,
  composto text,
  revisado text,
  dtarevisado text,
  interpretacao text,
  referencias jsonb
)
WITH (
  OIDS=FALSE
);
ALTER TABLE marcelo_result_final.informacoes_extra
  OWNER TO postgres;
