# M9 - Semana 5: Cypress
Este repositório contém os códigos e exemplos baseados no tutorial **"How to Test React using Cypress"**. Para os testes práticos, utilizei o site **Brazil Journal** como caso de estudo.

### Sumário
- Ponderada-Cypress
- Índice
  - Introdução
  - Configuração do Ambiente
  - Execução dos Testes
  - Documentação com Capturas de Tela
  - Conclusão

### Introdução
Este projeto foca na aplicação do **Cypress** para testar funcionalidades de uma aplicação web. Além disso, foram implementados testes personalizados no site dado no tutorial **https://reactjs.org/**, a fim de validar elementos críticos da página.

### Configuração do Ambiente
Para preparar o ambiente de desenvolvimento com Cypress e executar os testes, siga estas etapas:

1. **Inicializar um projeto Node.js**: Crie um arquivo `package.json` básico com o comando:
   ```bash
   npm init -y
   ```

2. **Instalar o Cypress**: Adicione o Cypress como uma dependência de desenvolvimento:
   ```bash
   npm install cypress --save-dev
   ```

3. **Abrir a interface do Cypress**: Após a instalação, você pode abrir a interface gráfica do Cypress com o seguinte comando:
   ```bash
   npx cypress open
   ```

4. **Estrutura de diretórios do Cypress**: Na primeira execução, o Cypress criará automaticamente os seguintes diretórios e arquivos:
   - **`cypress/`**: Contém as pastas `e2e`, `fixtures`, `support`, entre outras.
   - **`cypress.config.js`**: Arquivo de configuração para os testes.

### Criação de Testes
Navegue até o diretório `cypress/e2e` e crie um arquivo de teste, como por exemplo: `teste.cy.js`.

### Execução dos Testes

#### Testando o site 
No arquivo `teste.cy.js`, foram criados dois testes principais:
- O primeiro verifica se a página inicial é carregada corretamente.
- O segundo valida se os títulos das notícias estão visíveis na página.

### Documentação com Capturas de Tela

![Image](https://github.com/user-attachments/assets/340bb32e-b20a-4359-88e4-269adea38add)


![Image](https://github.com/user-attachments/assets/1c3de14f-3aad-4cf6-92ba-09ab4f4c9cab)

### Conclusão
Este repositório apresenta um guia prático para configurar e executar testes automatizados de interface com **Cypress**. Sendo uma ferramenta eficiente e simples de configurar, o Cypress permite simular interações de usuários em uma aplicação web, assegurando o correto funcionamento dos componentes. No contexto deste projeto, adaptamos os testes para validar o site, destacando a flexibilidade do Cypress em testar desde aplicações dinâmicas como aquelas feitas em React, até sites de conteúdo. Essa abordagem demonstra o poder do Cypress em diversos cenários de automação de testes, reforçando a prática de Desenvolvimento Orientado a Testes (TDD) como essencial para a entrega de software de alta qualidade.

---

Ao longo deste projeto, o conceito de TDD foi aplicado ao definir testes claros e verificáveis antes da implementação final, assegurando que cada componente ou funcionalidade do sistema fosse desenvolvido para atender a um comportamento específico pré-estabelecido nos testes.