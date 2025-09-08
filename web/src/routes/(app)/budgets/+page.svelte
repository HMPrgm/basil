<script>
    let budgets = [
        { id: 1, name: "Groceries", amount: 400 },
        { id: 2, name: "Rent", amount: 1200 },
        { id: 3, name: "Utilities", amount: 150 },
        { id: 4, name: "Entertainment", amount: 100 }
    ];
</script>

<h1>Budgets</h1>

{#if budgets.length === 0}
    <p>No budgets found.</p>
{:else}
    <table>
        <thead>
            <tr>
                <th>Name</th>
                <th>Amount</th>
            </tr>
        </thead>
        <tbody>
            {#each budgets as budget}
                <tr>
                    <td>{budget.name}</td>
                    <td>${budget.amount}</td>
                </tr>
            {/each}
        </tbody>
    </table>
{/if}