Ponderada de Programação: TDD com Golang

# Objetivo:
A partir do tutorial, deve-se executar cada exemplo, e documentar por prints as execuções. Ampliar os comentarios do código explicando as tecnicas e conceitos do TDD.

# TDD 
O **Desenvolvimento Orientado a Testes (TDD)** é uma prática de desenvolvimento de software que enfatiza a criação de testes automatizados antes da implementação do código funcional. Essa abordagem ajuda a garantir que o código atenda aos requisitos especificados e seja menos propenso a erros. O processo de TDD é geralmente dividido em três fases principais, conhecidas como o ciclo **Red-Green-Refactor**:

1. **Red (Vermelho)**: Nessa fase, o desenvolvedor escreve um teste automatizado para uma funcionalidade específica que ainda não foi implementada. O teste é executado e, como esperado, falha (fica "vermelho") porque a funcionalidade ainda não existe.

2. **Green (Verde)**: Depois que o teste falha, o próximo passo é escrever o código mínimo necessário para fazer o teste passar (ficar "verde"). Nesta fase, o foco está em implementar a funcionalidade de forma simples e direta.

3. **Refactor (Refatoração)**: Com o teste passando, o desenvolvedor pode refatorar o código para melhorar sua estrutura e eficiência, sem alterar o comportamento funcional. O teste garante que as melhorias no código não introduzam novos bugs.

# Configuração 
Antes de executar os testes, foi necessário a instalação da linguagem, a verificação da vesão no powershell e um teste simples de funcionamento, nesse repositório é possível encontrar o arquivo "main.go" para testar um print  

![Image](https://github.com/user-attachments/assets/0203ddcb-1a95-4516-ad8d-d5ac7d61112c)

![Image](https://github.com/user-attachments/assets/55a69ba0-b2b9-4f8f-9872-62be01958112)


# Aplicação do TDD no Código

Nos códigos fornecidos pelo tutorial, o TDD foi aplicado da seguinte maneira:

- **Red**: Primeiramente, testes foram escritos para as funções `SayHello`, `OddOrEven`, e `Checkhealth`. Esses testes especificavam os resultados esperados para cada função com base em diferentes entradas.

- **Green**: Em seguida, as funções foram implementadas para passar nos testes. Por exemplo, a função `SayHello` foi codificada para retornar uma mensagem de saudação personalizada, enquanto `OddOrEven` foi desenvolvida para verificar se um número é ímpar ou par.

- **Refactor**: Após os testes passarem, a refatoração poderia ser feita para melhorar a legibilidade e eficiência do código, mantendo todos os testes verdes. Esse processo garante que o código seja de alta qualidade e atenda aos requisitos definidos nos testes.

O uso de TDD nesses códigos assegura que as funções implementadas sejam rigorosamente verificadas quanto à sua precisão e comportamento esperado, promovendo um desenvolvimento mais robusto e confiável.

# Resultado dos testes 
Para executar os testes, foi o usado o comando: `go get github.com/stretchr/testify/assert` para instalar as dependências ncessárias e `go test -v` que faz com que os detalhes dos testes sejam exibidos, tendo assim, as saída indicando que todos os testes foram bem-sucedidos como indica as imagens a seguir 

![Image](https://github.com/user-attachments/assets/bfeac060-7e62-491a-9bcc-b5518cb31245)



![Image](https://github.com/user-attachments/assets/fc435ff5-fe00-49d0-8759-f6a51d822eb9)



![Image](https://github.com/user-attachments/assets/fba594c1-9373-4f11-b6d1-7249192d09b3)

