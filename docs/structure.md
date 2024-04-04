```mermaid
---
title: Projeto Integrador
---
---
title: Projeto Integrador
---
classDiagram
    namespace Cruds {
        class User{
            string id
            string name
            string email        
            string password
            string phone
            time birth_date
            RoleEnum role
        }

        class Pokemon{
            string id
            int reference_id
            string tier_id
            string name
            int weight
            int height
            string img_url 
            int experience
        }

        class Type{
            string id
            int reference_id
            string name
        }

        class Price{
            string pokemon_id
            time created_at
            int value
        }

        class Stock{
            string pokemon_id
            int quantity
        }
    }

    namespace Core {
        class Cart{
            string id
            string user_id
            time created_at
            time expires_in
            bool is_active
            int total
        }

        class Item{
            string id
            string cart_id
            string price_id
            string pokemon_id
            string purchase_id
            int quantity
            int total
        }

        class Purchase {
            string id
            string user_id
            int total
            bool is_approved
            time payment_limit_time
            string payment_method_id
        }

        class PaymentMethod {
            string id
            string name
        }
    }    

    class UserPokemon {
        string user_id
        string pokemon_id
        int quantity
    }

    namespace Diferencial {
        class UserStats{
            string user_id
            int experience
            string tier_id
        }


        class Tier{
            int id
            string name
            int minimal_experience
            int limit_experience
        }
    }
    

    Pokemon "1..N" --* "1..N" Type : possui
    Pokemon "1" --* "1..N" Price : tem
    Pokemon "1" --o "1" Stock : possui

    User "1" --o "0..N" Cart : possui
    Cart "1" *-- "1..N" Item : é composto
    Item "1" *-- "1" Price : possuí
    Pokemon "1" *-- Item : refere-se a
    Purchase "1" --* "1..N" Item : é composta de
    Purchase "0..N" --* "1" User : possuí
    Purchase "0..N" --o "1" PaymentMethod : possuí

    User "1" --> "N" UserPokemon 
    Pokemon "1" --> "N" UserPokemon 

    User "1" *-- "1" UserStats : possuí  

    UserStats "0..N" --* "1" Tier : está 
    Pokemon "0..N" --* "1" Tier : está 
```