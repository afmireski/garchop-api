# Changelog

## [Unreleased]

* [#75](https://github.com/afmireski/garchop-api/issues/75) - Implementar rota que atualiza a quantidade de um item no carrinho
* [#60](https://github.com/afmireski/garchop-api/issues/60) - Implementar exclusão de recompensa
* [#71](https://github.com/afmireski/garchop-api/issues/71) - Implementar listagem específica de Pokémons de um usuário

## [Release 2](https://github.com/afmireski/garchop-api/releases/tag/24.08.23)

### Added
* [#61](https://github.com/afmireski/garchop-api/issues/61) - Implementar obtenção de recompensas por usuários
    * Removendo parâmetros `user_id` das rotas e substituindo por obtenção via `token`
    * Passagem de parâmetro `filtro` na query da rota `/pokemon`
    * Implementada rota para listagem de formas de pagamento
* [#58](https://github.com/afmireski/garchop-api/issues/56) - Implementar listagem geral de recompensas
* [#57](https://github.com/afmireski/garchop-api/issues/57) - Implementar cadastro de recompensas
* [#52](https://github.com/afmireski/garchop-api/issues/52) - Obter Status do Usuário
* [#47](https://github.com/afmireski/garchop-api/issues/47) - Implementar autenticação
* [#58](https://github.com/afmireski/garchop-api/issues/58) - Implementar listagem geral de recompensas
* [#51](https://github.com/afmireski/garchop-api/issues/50) - Atualizar a experiência de usuário ao comprar Pokémons
* [#56](https://github.com/afmireski/garchop-api/issues/56) - Trazer dados de Status do Usuário na hora de listar o perfil do usuário
* [#45](https://github.com/afmireski/garchop-api/issues/45) - Implementar listagem do histórico de compras do usuário
* [#50](https://github.com/afmireski/garchop-api/issues/50) - Criar `UserStats` no momento do cadastro do usuário
* [#39](https://github.com/afmireski/garchop-api/issues/39) - Implementar remoção de um item do carrinho de compras
* [#37](https://github.com/afmireski/garchop-api/issues/37) - Implementar busca do carrinho de compra de um usuário
    * Definição das portas de ItemsRepository e CartsRepository
    * Implementação dos adaptadores do Supabase de ItemsRepository e CartsRepository
    * Implementação da rota: `/users/:user_id/cart`

### Changed

### Fixed
* [#64](https://github.com/afmireski/garchop-api/issues/64) - [HOTFIX] Corrigir lógica de ganho de experiência
* [#68](https://github.com/afmireski/garchop-api/issues/68) - [BUG]: Informações detalhadas de um pokémon não aparecem na compra
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
