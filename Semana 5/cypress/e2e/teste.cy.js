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

describe('Navigation Test', () => {
  it('should navigate to the About page', () => {
    cy.visit('http://localhost:3000');

    // Encontra e clica no link para a página "About"
    cy.get('a[href*="about"]').click();

    // Verifica se o conteúdo da página About é exibido
    cy.url().should('include', '/about');
    cy.contains('About Us');
  });
});
