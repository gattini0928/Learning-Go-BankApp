<!-- Usuário possui uma conta -->
<!-- Criar conta -->
<!-- Validar dados -->
<!-- Email + Nome + CPF + Senhas(hashar senhas) -->
<!-- Criar user com dados insert no banco de dados --> 
<!-- Accounts db -> tabela users -->
<!-- Funcs de Criar cartão de crédito -->
<!-- Depósito, transferências, compra parcelada, Upar de nível de conta
Acessar Fatura do cartão de crédito, Pedir Empréstimo,Logout, Excluir Conta -->
<!-- Contas -> Silver, Gold, Platinum, Diamond, Premium Diamond -->
<!-- Bronze: Gratuita default -->
<!-- Silver: $30 por mes -->
<!-- Gold: $50 por mes -->
<!-- Platinum: $70 por mes -->
<!-- Diamond: $120 por mes -->
<!-- Premium Diamond: $200 por mes -->
<!-- Login (auth credencias) -->

bool -> signup -> config senha entrar + senha bancaria -> true

if true -> for loop -> Transferir, Pagar Cartão de Crédito, Ver Faturas, Consultar saldo

Crédito ou Débito -> Débito(Tirar saldo), Crédito(Adicionar na fatura sem tirar o saldo)

SE a compra no crédito for maior que o limite travar
Sempre que comprar algo - diminuir limite float64 Ex: 200.00 -> -200.00 de 400.00(default)

Compra parcelada baseada num valor(até 20 reais (1x), > 20 < 60(2x), > 60 < 100 (4x), > 100 (12x)) - Total das parcelas diminur Ex: 400.00 - 200.00 feito em 1x,2x,4x,12x

map[string]float64 -> Produtos para testar valores e erros

Comunicação entre contas -> Transferir(+ pra quem ganhou - pra quem transferir)


