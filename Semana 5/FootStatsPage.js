class FootStatsPage {
    // Seletores da página
    elements = {
      menuButton: () => cy.xpath('//*[@id="menu-itens"]/li[1]/a'), // Exemplo de  menu
      mainContent: () => cy.get('.main-content'), // Exemplo de conteúdo principal
      footer: () => cy.get('footer'), // Selecionando o rodapé
      specificLink: () => cy.get('.verDetalhes') // Link específico na SPA
    }
  
    // Métodos para interagir com a página
    clickMenuButton() {
      this.elements.menuButton().click();
    }
  
    navigateToSpecificSection() {
      this.elements.specificLink().click();
    }
  
    checkMainContentLoaded() {
      this.elements.mainContent().should('be.visible');
    }
  
    checkFooterVisible() {
      this.elements.footer().should('be.visible');
    }
  }
  
  module.exports = new FootStatsPage();