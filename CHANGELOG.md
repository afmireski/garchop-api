# Changelog

## [Unreleased]

### Added
* [#39](https://github.com/afmireski/garchop-api/issues/39) - Implementar remoção de um item do carrinho de compras
* [#37](https://github.com/afmireski/garchop-api/issues/37) - Implementar busca do carrinho de compra de um usuário
    * Definição das portas de ItemsRepository e CartsRepository
    * Implementação dos adaptadores do Supabase de ItemsRepository e CartsRepository
    * Implementação da rota: `/users/:user_id/cart`

### Changed

### Fixed
* [#34](https://github.com/afmireski/garchop-api/issues/34) - BUG: Usuários excluídos estão sendo listados nas buscas de usuário

## [Release 1](https://github.com/afmireski/garchop-api/releases/tag/v1.0.0)

### Added
* [#21](https://github.com/afmireski/garchop-api/issues/27) - Implementar rota que permita a um administrador editar os detalhes de um Pokémon
* [#13](https://github.com/afmireski/garchop-api/issues/13) - Implementar rota que permita a listagem de usuários administrativos
* [#27](https://github.com/afmireski/garchop-api/issues/27) - Implementar listagem dos Tiers existentes
* [#10](https://github.com/afmireski/garchop-api/issues/10) - Implementar rota de cadastro de conta administrativa
* [#11](https://github.com/afmireski/garchop-api/issues/11) - Implementar rota que permita retirar um pokémon de venda
* [#12](https://github.com/afmireski/garchop-api/issues/12) - Implementar rota que permita listar todos os pokémons que estão a venda
* [#20](https://github.com/afmireski/garchop-api/issues/20) - Implementar rota que exiba as informações detalhadas de um Pokémon
* [#9](https://github.com/afmireski/garchop-api/issues/9) - Implementar rota que permita o cadastro de um novo pokémon
    * Criado novos recursos relacionados a pokemons e tipo de pokemon
* [#4](https://github.com/afmireski/garchop-api/issues/4) - Implementar rota que permita a atualização das informações de um usuário
* [#16](https://github.com/afmireski/garchop-api/issues/16) - Implementar rota de Login na plataforma
    * Refatoração do cadastro de usuário para registrar usuário na tabela de autenticação do supabase.
* [#5](https://github.com/afmireski/garchop-api/issues/5) - Implementar rota para excluir um usuário
* [#3](https://github.com/afmireski/garchop-api/issues/3) - Implementar rota que possibilite o cadastro de um usuário
* [#2](https://github.com/afmireski/garchop-api/issues/2) - Implementar rota para buscar informações de um usuário

### Changed

### Fixed