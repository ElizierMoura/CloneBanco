package models

import "github.com/eduardo-mior/gorm/dialects/postgres"

type ExameDetalhado struct {
	Status string   `json:"status" gorm:"-"`
	Exames []Exames `json:"exames"`
}

type Exames struct {
	Exame Exame `json:"exame"`
}

type Exame struct {
	Mneumonico       string           `json:"mneumonico"`
	Nome             string           `json:"nome"`
	Material         string           `json:"material"`
	PreRequisitos    string           `json:"preRequisitos"   gorm:"column:prerequisitos"`
	Ativo            bool             `json:"ativo"`
	Unidades         string           `json:"unidades"`
	InformacoesExtra InformacoesExtra `json:"informacoesExtra"`
}

type InformacoesExtra struct {
	ID                          interface{}    `json:"id"`
	Codigo                      interface{}    `json:"codigo"`
	Nome                        interface{}    `json:"nome"`
	Material                    interface{}    `json:"material"`
	Sinonimo                    interface{}    `json:"sinonimo"`
	Volume                      interface{}    `json:"volume"`
	Metodo                      interface{}    `json:"metodo"`
	VolumeLab                   interface{}    `json:"volume_lab"`
	VolumeMinimo                interface{}    `json:"volume_minimo"`
	EstabilidadeAmbienteQtd     interface{}    `json:"estabilidade_ambiente_qtd"`
	EstabilidadeAmbienteTipo    interface{}    `json:"estabilidade_ambiente_tipo"`
	EstabilidadeRefrigeradaQtd  interface{}    `json:"estabilidade_refrigerada_qtd"`
	EstabilidadeRefrigeradaTipo interface{}    `json:"estabilidade_refrigerada_tipo"`
	EstabilidadeFreezerQtd      interface{}    `json:"estabilidade_freezer_qtd"`
	EstabilidadeFreezerTipo     interface{}    `json:"estabilidade_freezer_tipo"`
	Rotina                      interface{}    `json:"rotina"`
	Resultado                   interface{}    `json:"resultado"`
	Temperatura                 interface{}    `json:"temperatura"`
	Coleta                      interface{}    `json:"coleta"`
	CodigoSus                   interface{}    `json:"codigo_sus"`
	Cbhpm                       interface{}    `json:"cbhpm"`
	Interpretacao               interface{}    `json:"interpretacao"`
	Esoterico                   interface{}    `json:"esoterico"`
	Composto                    interface{}    `json:"composto"`
	Revisado                    interface{}    `json:"revisado"`
	DtaRevisado                 interface{}    `json:"dta_revisado"`
	Referencias                 postgres.Jsonb `json:"referencias"`
}

type ref struct {
	Referencia interface{} `json:"referencia"`
}
