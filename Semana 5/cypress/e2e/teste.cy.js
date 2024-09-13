describe('React App Load Test', () => {
  it('should load the React app', () => {
    // Visita a aplicação em execução
    cy.visit('https://reactjs.org/');

    // Verifica se o título da página contém o texto correto
    cy.title().should('include', 'React App');

    // Verifica se o conteúdo principal é carregado
    cy.contains('Welcome to React');
  });
});
