package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//resposta define a estrutura do json que será retornado
type Resposta struct {
	Nome string `json:"nome"`
	Horario string `json:"horario"`
}

//totalRequisicoes conta as vezes que o endpoint foi chamado
var totalRequisicoes = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "http_requisicoes_total",
	Help: "Total de requisições recebidas no endpoint /projeto-korp",
})

//servicoDisponivel indica se o serviço ainda ta on
var servicoDisponivel = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "servico_disponivel",
	Help: "Disponibilidade do serviço: 1 = disponível, 0 = indisponível",
})

//endpointProjetoKorp é o handler do get /projeto-korp
func endpointProjetoKorp(w http.ResponseWriter, r *http.Request) {
	//incrementa o contador a cada requisição
	totalRequisicoes.Inc()

	//pega o horario de agora em utc
	horarioAtual := time.Now().UTC().Format(time.RFC3339)

	resposta := Resposta{
		Nome: "Projeto Korp",
		Horario: horarioAtual,
	}

	//isso faz o client reconhecer que está recebendo um json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resposta)
}

func main() {
	//registra as metricas no prometheus
	prometheus.MustRegister(totalRequisicoes)
	prometheus.MustRegister(servicoDisponivel)

	//marca o serviço como disponivel quando subir
	servicoDisponivel.Set(1)
	

	//endpoint do app
	http.HandleFunc("/projeto-korp", endpointProjetoKorp)

	//endpoint de metricas
	http.Handle("/metrics", promhttp.Handler())	

	//logs
	log.Println("Servidor rodando na porta 8080")
	log.Println("Métricas disponíveis em /metrics")
	log.Fatal(http.ListenAndServe(":8080", nil))
}