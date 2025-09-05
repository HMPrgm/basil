<script>
    let username = "";
    let email = "";
    let password = "";

    function handleSubmit() {
        console.log("Registering:", { username, email, password });
    }
</script>


<main>
    <div class="container mx-auto mt-10 p-4">
        <form on:submit|preventDefault={handleSubmit} class="flex flex-col gap-4 max-w-md mx-auto mt-8 bg-slate-50 p-6 rounded-lg shadow-md">
            <h2 class="text-2xl font-bold mb-4 text-center">Create an Account</h2>

            <label for="username">Username:</label>
            <input type="text" id="username" name="username" bind:value={username} class="border p-2 rounded-md focus:outline-none focus:border-emerald-500" required>
            <label for="email">Email:</label>
            <input type="email" id="email" name="email" bind:value={email} class="border p-2 rounded-md focus:outline-none focus:border-emerald-500" required>
            <label for="password">Password:</label>
            <input type="password" id="password" name="password" bind:value={password} class="border p-2 rounded-md focus:outline-none focus:border-emerald-500" required>
            <button type="submit" class="bg-emerald-500 text-white p-2 rounded-md hover:bg-emerald-600 transition-colors">Register</button>
        </form>
    </div>
</main>