Registro consolidado das decisões técnicas do projeto.
Este documento concentra todas as decisões técnicas relevantes tomadas ao longo do desenvolvimento.  
Cada decisão deve ser registrada em ordem cronológica, identificada por um número sequencial e acompanhada de sua data.

---
## 📌 Formato para novas decisões

Use este modelo sempre que for adicionar uma nova decisão:

### **TDR-XXX — Título da decisão**  
**Data:** YYYY-MM-DD  

**Contexto:**  
Descreva brevemente o cenário que originou a necessidade da decisão.

**Problema Identificado:**  
Explique o problema técnico que precisava ser resolvido.

**Decisão:**  
Registre claramente a decisão tomada.

**Justificativa:**  
Apresente os motivos que levaram à escolha.

**Consequências:**  
Liste os impactos positivos e negativos.

---

## 📚 Decisões Registradas

---

### **TDR-002 — Consolidação dos dados de transações em relatórios**  
**Data:** 2025-11-16  

**Contexto:**  
O projeto contém uma feature de relatório gerais baseado nas somas das entradas ,saídas e aportes que o usuário cadastrou

**Problema Identificado:**  
O problema que identifique foi a onde e quando eu iria processar esses dados , no schema de transaction (com a criação de diversas views), no micro serviço de transactions (com end-point para cada relatório)

**Decisão:**  
Decidi então criar um schema direto para a consolidação desses dados seguindo a seguinte logica da planilha de base do sistema , onde irei registrar as somas de cada tipo de transaction
![[Pasted image 20251116075218.png]]

**Justificativa:**  
- Facildade na hora de consultar os relatórios
- Menos joins em tabelas 
- Menor processamento na hora de retornar os relatórios 

**Consequências:**  
 - Positivas: 
	 - consolidação de dados centralizados
	 - relatórios centralizados
 - Negativas: 
	 - um schema a mais para manter
	 - envio de eventos obrigatório a cada mudança em qualquer transaction e investment

---


### **TDR-001 — Reorganização dos pacotes `middleware` e `pkg` para evitar dependência circular**  
**Data:** 2025-11-13  

**Contexto:**  
No início do projeto, a arquitetura foi planejada com módulos separados por conceito: cada micros-serviço possui seu próprio módulo, enquanto pacotes utilitários ficaram em um módulo `pkg`, e os *middlewares* em um módulo dedicado. Os *middlewares* são responsáveis por processamentos intermediários das rotas, como validação de autenticação e registro de logs.

**Problema Identificado:**  
Durante o desenvolvimento, alguns *middlewares* passaram a depender de estruturas definidas no módulo `pkg`, especialmente o struct `Response`. Isso criava risco de duplicação de código e potencial dependência circular entre os módulos.

**Decisão:**  
O pacote `middleware` foi movido para dentro do mesmo módulo do `pkg`, e o módulo antigo de *middlewares* foi removido.

**Justificativa:**  
A unificação elimina a possibilidade de dependência circular, evita a reescrita do struct `Response`, centraliza os tipos compartilhados e reduz a complexidade estrutural do projeto. Além disso, mantém os *middlewares* próximos dos recursos comuns que realmente utilizam.

**Consequências:**  
- **Positivas:** arquitetura mais simples, build sem conflitos, código centralizado e sem duplicações.  
- **Negativas:** menor isolamento entre pacotes, podendo exigir ajustes futuros caso o módulo de *middlewares* precise ser extraído.

---

