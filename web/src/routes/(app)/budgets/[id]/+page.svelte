<script>
    // Example budgets data; replace with your own data source or fetch from an API
    let budget = { id: 1, name: "Groceries", amount: 400 };
</script>

<h1>{budget.name}</h1>
    <p>Amount: ${budget.amount}</p>
    <p>Spent: ${budget.amount/2}</p>
    <p>Remaining: ${budget.amount - (budget.amount/2)}</p>