# desafio-multithreading

Desafio feito para o curso de Go Expert da plataforma FullCycle

## Objetivos

Esse desafio consiste em fazer duas chamadas a APIs distintas de maneira simultanea e retornar o resultado da api que retornará o resultado mais rapidamente.
As Apis em questão são as <a href="https://viacep.com.br">viacep</a> e a <a href="https://brasilapi.com.br">brasilapi</a>.

## Requisitos
- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.