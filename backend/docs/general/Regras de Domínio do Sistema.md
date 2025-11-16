
---

# 📘 **REGRAS DE NEGÓCIO DO Gofre (Baseado na Planilha)**

---

# 1. **TRANSACTIONS SERVICE**

(Agrupa ENTRADA e SAÍDA — mov. financeiras)

## 1.1. Objetivo

Registrar todas as movimentações financeiras do usuário e disponibilizar cálculos consolidados por mês.

## 1.2. Regras de negócio

1. Toda transação deve ter:
    
    - data
        
    - categoria
        
    - descrição
        
    - valor
        
    - tipo (`entrada` ou `saida`)
        
2. O valor deve ser **sempre positivo**.
    
3. A categoria deve existir na lista fixa (front) ou no service `categories` (caso exista).
    
4. O sistema deve agrupar transações por:
    
    - mês
        
    - categoria
        
    - tipo
        
5. Resultado mensal das transações:
    
    - **entradas_totais = soma(entradas do mês)**
        
    - **saidas_totais = soma(saídas do mês)**
        
    - **saldo = entradas_totais – saidas_totais**
        
6. Este serviço deve fornecer totais para outros microserviços:
    
    - goals → comparação planejado × realizado
        
    - forecast → projeções
        
    - reports → dashboard
        
    - property → impacto no patrimônio (opcional)
        

---

# 2. **INVESTMENTS SERVICE**

(Representa o que na planilha é APORTE)

## 2.1. Objetivo

Registrar aportes e calcular o total aportado no mês/ano.

## 2.2. Regras de negócio

1. Um aporte deve ter:
    
    - data
        
    - local (ex: corretora, banco)
        
    - descrição
        
    - valor
        
2. O valor deve ser positivo.
    
3. A soma mensal de aportes deve ser calculada.
    
4. A soma anual deve ser acumulada.
    
5. Os aportes alimentam:
    
    - reports (resultado mensal)
        
    - property (patrimônio)
        
    - forecast (previsão)
        
6. Não são permitidas datas futuras.
    

---

# 3. **GOALS SERVICE**

(Representa MENSAL → Metas financeiras)

## 3.1. Objetivo

Gerenciar metas planejadas para o mês: entrada, saída, aporte e resultado desejado.

## 3.2. Regras de negócio

1. O usuário pode definir metas:
    
    - meta_entrada
        
    - meta_saida
        
    - meta_aporte
        
    - saldo_minimo_desejado
        
    - resultado_esperado
        
2. Uma meta pode ser alterada apenas para o **mês atual** ou para meses futuros.
    
3. O sistema deve calcular:
    
    - diferença = realizado – meta
        
    - percentual_atingido = (realizado/meta) × 100
        
4. Deve comparar os valores realizados vindos de:
    
    - transactions
        
    - investments
        
5. As metas alimentam:
    
    - reports (planejado vs realizado)
        
    - forecast (projeções futuras)
        

---

# 4. **FORECAST SERVICE**

(Representa PREVISAO)

## 4.1. Objetivo

Gerar projeções futuras de entradas, saídas, aportes e saldo.

## 4.2. Regras de negócio

1. A previsão deve utilizar:
    
    - média dos últimos meses
        
    - metas do mês atual (goals)
        
    - valores recorrentes
        
2. Se uma movimentação é marcada como recorrente:
    
    - ela deve ser replicada automaticamente para cada mês futuro.
        
3. Deve calcular:
    
    - previsão de entrada mensal
        
    - previsão de saída mensal
        
    - previsão de aporte
        
    - saldo previsto
        
    - resultado previsto
        
4. A previsão deve ser recalculada quando:
    
    - transações mudarem
        
    - aportes mudarem
        
    - metas mudarem
        
5. A previsão deve fornecer:
    
    - projeção dos próximos 12 meses
        
    - projeção anual consolidada
        

---

# 5. **PROPERTY SERVICE**

(Representa PATRIMONIO)

## 5.1. Objetivo

Controlar a evolução do patrimônio total do usuário.

## 5.2. Regras de negócio

1. Patrimônio total deve considerar:
    
    - aportes acumulados
        
    - valorização dos ativos
        
    - patrimônio inicial registrado pelo usuário
        
2. Cálculo básico:
    
    ```
    patrimonio_total = patrimonio_anterior + aporte_mensal + retorno
    ```
    
3. O usuário pode registrar patrimônio manual:
    
    - bens
        
    - investimentos fora do sistema
        
4. Deve calcular:
    
    - crescimento absoluto
        
    - crescimento percentual
        
5. Deve fornecer dados para o reports e simulator.
    

---

# 6. **SIMULATOR SERVICE**

(Representa SIMULADOR → juros compostos)

## 6.1. Objetivo

Calcular juros compostos e previsões financeiras.

## 6.2. Regras de negócio

### Cálculo base:

```
montante = P * (1 + i)^n + A * [((1+i)^n - 1) / i]
```

Onde:

- P = valor inicial
    
- i = taxa de juros mensal
    
- A = aporte mensal
    
- n = número de meses
    

### Regras:

1. O simulador deve permitir:
    
    - **simulação por objetivo**: calcular aporte necessário
        
    - **simulação por tempo**: calcular meses para atingir objetivo
        
2. Deve validar:
    
    - taxa > 0
        
    - aporte ≥ 0
        
    - valor inicial ≥ 0
        
3. O simulador deve gerar:
    
    - total investido
        
    - total em juros
        
    - montante final
        
    - meses necessários
        
4. Pode opcionalmente enviar seu resultado para:
    
    - property (atualizar patrimônio real)
        

---

# 7. **REPORTS SERVICE**

(Representa o dashboard RESULTADO)

## 7.1. Objetivo

Consolidar informações de todos os serviços e gerar a visão geral do mês.

## 7.2. Regras de negócio

O relatório mensal deve conter:

### Dados consolidados

1. **Entrada total** (transactions)
    
2. **Saída total** (transactions)
    
3. **Aporte total** (investments)
    
4. **Saldo** = entrada – saída
    
5. **% aporte** em relação à meta (goals)
    
6. **Resultado do mês**
    
7. **Planejado × Realizado**
    
8. **Patrimônio total** (property)
    
9. **Previsão dos próximos meses** (forecast)
    

### Regras

1. O reports faz **somente leitura**, sem gravar nada.
    
2. Ele deve consumir dados dos outros microserviços.
    
3. O cálculo deve ser sempre em **tempo real**.
    
4. Deve fornecer indicadores:
    
    - variação de mês anterior
        
    - diferença entre metas e valores reais
        
    - percentual de atingimento por categoria
        
5. Deve exibir alertas:
    
    - saldo negativo
        
    - gasto maior que meta
        
    - aporte menor que meta
        
    - patrimônio caindo
        

---

# 8. **CATEGORIES (opcional)**


### Regras (caso exista):

1. CRUD de categorias.
    
2. Não pode haver categoria duplicada.
    
3. Categoria não pode ser excluída caso exista transação vinculada.
    
4. Tipos de categoria:
    
    - entrada
        
    - saída
        
    - aporte (opcional)
        
5. As categorias alimentam transactions/investments.
    

---
