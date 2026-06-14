# korp-devopsjr-desafio

Desafio técnico para a vaga de Analista de DevOps Júnior / Estágio em DevOps da Korp ERP.

Servidor HTTP em GO com proxy reverso nginx, monitoramento via Prometheus junto com o Grafana e provisionamento automatizado com Ansible.

---

## Tecnologias que foram utilizadas no projeto

- Golang, Docker, Docker compose, Nginx, Prometheus, Grafana e Ansible

---

## Estrutura do projeto

```
korp-devopsjr-desafio/
├── app/                  # servidor Go + Dockerfile
├── nginx/                # configuração do proxy reverso
├── prometheus/           # configuração de scraping
├── compose.yaml          # orquestração dos containers
└── playbook.yml          # automação com Ansible
```

---

## Como rodar

Necessário: Linux ou wsl2, ansible e o community.docker

```bash
# Instalar dependências do Ansible
ansible-galaxy collection install community.docker

# Provisionar tudo com um comando
ansible-playbook playbook.yml
```

O playbook instala o docker, faz o build da imagem, sobe os containers e valida o serviço tudo automaticamente

---

## Endpoints

| Serviço | Endereço |
|---|---|
| Aplicação | http://localhost:80/projeto-korp |
| Prometheus | http://localhost:9090 |
| Grafana | http://localhost:3000 |

Grafana: user: `admin` / senha: `admin123`

---

## Métricas

O dashboard http-server-projeto-korp no Grafana exibe disponibilidade do serviço e volume de requisições