package controled

import (
	"apiGetBanco/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Estrutura de Exame/s
type examesListagem struct {
	Exames []Exames `json:"exames"`
}

type Exames struct {
	Exame Exame `json:"exame"`
}
type Exame struct {
	Mneumonico string `json:"mneumonico"`
}

func GetBanco(w http.ResponseWriter, r *http.Request) {

	// colorReset := "\033[0m"

	colorRed := "\033[31m"
	colorGreen := "\033[32m"
	colorYellow := "\033[33m"
	colorBlue := "\033[34m"

	fmt.Println(string(colorBlue), "Iniciando Função do GetBanco")

	dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	alfabeto := [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	for _, letra := range alfabeto {
		fmt.Println(string(colorYellow), "Iniciamos com a letra ", letra)
		fmt.Println("-------------------")
		fmt.Println("-------------------")
		fmt.Println("-------------------")
		fmt.Println("-------------------")
		examesList := retornarExames(letra)

		for _, v := range examesList.Exames {
			fmt.Println("-------------------")
			fmt.Println("-------------------")
			fmt.Println("-------------------")
			fmt.Println(string(colorYellow), "INICIANDO CADASTRAMENTO DO ", v.Exame.Mneumonico)
			exameDetalhado := exameDetalhado(v.Exame.Mneumonico)

			cadastrar := db.
				Table("marcelo_result_final.exames").
				Create(map[string]interface{}{
					"mneumonico": exameDetalhado.Exames[0].Exame.Mneumonico,
					"nome":       exameDetalhado.Exames[0].Exame.Nome,
					"material":   exameDetalhado.Exames[0].Exame.Material,
					"ativo":      exameDetalhado.Exames[0].Exame.Ativo,
					"unidades":   exameDetalhado.Exames[0].Exame.Unidades,
				})

			if cadastrar.Error != nil {
				fmt.Println(cadastrar.Error)
				fmt.Println(string(colorRed), "erro", v.Exame.Mneumonico)
				erro := db.
					Table("marcelo_result_final.erro").
					Create(map[string]interface{}{"mneumonico": exameDetalhado.Exames[0].Exame.Mneumonico})

				if erro != nil {
					fmt.Println(string(colorRed), "erro ao guardar erro")
				}
			}

			fmt.Println(exameDetalhado.Exames[0].Exame.InformacoesExtra.Referencias)

			cadastrar = db.
				Table("marcelo_result_final.informacoes_extra").
				Create(map[string]interface{}{
					"id":                          exameDetalhado.Exames[0].Exame.InformacoesExtra.ID,
					"codigo":                      exameDetalhado.Exames[0].Exame.InformacoesExtra.Codigo,
					"nome":                        exameDetalhado.Exames[0].Exame.InformacoesExtra.Nome,
					"material":                    exameDetalhado.Exames[0].Exame.InformacoesExtra.Material,
					"sinonimo":                    exameDetalhado.Exames[0].Exame.InformacoesExtra.Sinonimo,
					"volume":                      exameDetalhado.Exames[0].Exame.InformacoesExtra.Volume,
					"metodo":                      exameDetalhado.Exames[0].Exame.InformacoesExtra.Metodo,
					"volumelab":                   exameDetalhado.Exames[0].Exame.InformacoesExtra.VolumeLab,
					"volumeminimo":                exameDetalhado.Exames[0].Exame.InformacoesExtra.VolumeMinimo,
					"estabiliadeambienteqtd":      exameDetalhado.Exames[0].Exame.InformacoesExtra.EstabilidadeAmbienteQtd,
					"estabilidadeambientetipo":    exameDetalhado.Exames[0].Exame.InformacoesExtra.EstabilidadeAmbienteTipo,
					"estabilidaderefrigeradaqtd":  exameDetalhado.Exames[0].Exame.InformacoesExtra.EstabilidadeRefrigeradaQtd,
					"estabilidaderefrigeradatipo": exameDetalhado.Exames[0].Exame.InformacoesExtra.EstabilidadeRefrigeradaTipo,
					"estabilidadefreezerqtd":      exameDetalhado.Exames[0].Exame.InformacoesExtra.EstabilidadeFreezerQtd,
					"estabilidadefreezertipo":     exameDetalhado.Exames[0].Exame.InformacoesExtra.EstabilidadeFreezerTipo,
					"rotina":                      exameDetalhado.Exames[0].Exame.InformacoesExtra.Rotina,
					"resultado":                   exameDetalhado.Exames[0].Exame.InformacoesExtra.Resultado,
					"temperatura":                 exameDetalhado.Exames[0].Exame.InformacoesExtra.Temperatura,
					"coleta":                      exameDetalhado.Exames[0].Exame.InformacoesExtra.Coleta,
					"codigosus":                   exameDetalhado.Exames[0].Exame.InformacoesExtra.CodigoSus,
					"cbhpm":                       exameDetalhado.Exames[0].Exame.InformacoesExtra.Cbhpm,
					"interpretacao":               exameDetalhado.Exames[0].Exame.InformacoesExtra.Interpretacao,
					"esoterico":                   exameDetalhado.Exames[0].Exame.InformacoesExtra.Esoterico,
					"composto":                    exameDetalhado.Exames[0].Exame.InformacoesExtra.Composto,
					"revisado":                    exameDetalhado.Exames[0].Exame.InformacoesExtra.Revisado,
					"dtarevisado":                 exameDetalhado.Exames[0].Exame.InformacoesExtra.DtaRevisado,
					"referencias":                 exameDetalhado.Exames[0].Exame.InformacoesExtra.Referencias,
				})

			if cadastrar.Error != nil {
				fmt.Println(cadastrar.Error)
				fmt.Println(string(colorRed), "erro :", v.Exame.Mneumonico)
				erro := db.
					Table("marcelo_result_final.erro").
					Create(map[string]interface{}{"mneumonico": exameDetalhado.Exames[0].Exame.Mneumonico})

				if erro != nil {
					fmt.Println(string(colorRed), "erro ao guardar erro")
				}
			}

			fmt.Println(string(colorGreen), "TERMINANDO CADASTRAMENTO DO ", v.Exame.Mneumonico)
			fmt.Println("-------------------")
			fmt.Println("-------------------")
			fmt.Println("-------------------")
		}
		fmt.Println("-------------------")
		fmt.Println("-------------------")
		fmt.Println("-------------------")
		fmt.Println(string(colorBlue), "FIM letra: ", letra)
	}

	JSON(w, 200, "sucesso")

}

func exameDetalhado(idNomeExame string) models.ExameDetalhado {

	resp, err := http.Get(fmt.Sprintf("https://alvaroapoio.com.br/dataprovider_middleware/exames/detalhes-exame/%s?_format=json", idNomeExame))
	if err != nil {
		fmt.Println(string("\033[31m"), "erro")
		dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
		db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		erro := db.
			Table("marcelo_banco.erro").
			Create(map[string]interface{}{"mneumonico": idNomeExame})

		if erro != nil {
			fmt.Println("erro ao guardar erro")
		}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("erro")
		dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
		db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		erro := db.
			Table("marcelo_banco.erro").
			Create(map[string]interface{}{"mneumonico": idNomeExame})

		if erro != nil {
			fmt.Println("erro ao guardar erro")
		}
	}

	var exls models.ExameDetalhado

	if err := json.Unmarshal(body, &exls); err != nil {
		fmt.Println(string("\033[31m"), "erro")
		dsn := "host=localhost user=postgres password=1234 dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
		db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		erro := db.
			Table("marcelo_banco.erro").
			Create(map[string]interface{}{"mneumonico": idNomeExame})

		if erro != nil {
			fmt.Println("erro ao guardar erro")
		}
	}
	fmt.Println("chegou aqui com sucesso")
	return exls

}

func retornarExames(letra string) examesListagem {
	resp, err := http.Get(fmt.Sprintf("https://alvaroapoio.com.br/dataprovider_middleware/exames/%s?_format=json&rm=first-letter", letra))
	if err != nil {
		fmt.Println(string("\033[31m"), "erro")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("erro")

	}

	var exls examesListagem

	if err := json.Unmarshal(body, &exls); err != nil {
		fmt.Println(string("\033[31m"), "erro")
	}

	return exls
}

func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	fmt.Println("chegout aqui")
	if err := json.NewEncoder(w).Encode(dados); err != nil {
		fmt.Println("erro")
		fmt.Println(string("\033[31m"), "erro")
	}

}
