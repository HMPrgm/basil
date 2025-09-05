<svelte:head>
  <title>Basil - Your Financial Super App</title>
  <meta
    name="description"
    content="Basil is the all-in-one app to track spending, manage budgets, and achieve your financial goals effortlessly."
  />
</svelte:head>

<!-- Header -->
<header class="bg-gray-50/80 backdrop-blur-lg fixed top-0 left-0 right-0 z-50">
  <nav class="container mx-auto px-6 py-4 flex justify-between items-center">
    <div class="text-2xl font-bold text-emerald-600">Basil</div>
    <div class="flex items-center space-x-4">
      <a href="/login" class="text-gray-600 hover:text-emerald-600 font-medium"
        >Log In</a
      >
      <a
        href="/register"
        class="bg-emerald-600 text-white px-4 py-2 rounded-lg font-semibold hover:bg-emerald-700 transition-colors"
      >
        Sign Up Free
      </a>
    </div>
  </nav>
</header>

<main>
  <!-- Hero Section -->
  <section class="pt-32 pb-16 md:pt-48 md:pb-24">
    <div class="container mx-auto px-6 text-center">
      <h1
        class="text-4xl md:text-6xl font-extrabold tracking-tight text-gray-900 leading-tight pb-8"
      >
        Gain control of your money. <br />
        <span class="text-emerald-600">Effortlessly.</span>
      </h1>

      <div class="grid md:grid-cols-3 gap-8">
        <!-- Feature 1 -->
        <div class="text-center p-8 bg-gray-50 rounded-xl">
          <div
            class="bg-emerald-100 text-emerald-600 w-16 h-16 rounded-full mx-auto flex items-center justify-center"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-8 w-8"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              ><path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M13 10V3L4 14h7v7l9-11h-7z"
              /></svg
            >
          </div>
          <h3 class="mt-6 text-xl font-semibold text-gray-900">
            Auto-Track Spending
          </h3>
          <p class="mt-2 text-gray-600">
            Securely link your accounts and we'll automatically categorize your
            transactions.
          </p>
        </div>
        <!-- Feature 2 -->
        <div class="text-center p-8 bg-gray-50 rounded-xl">
          <div
            class="bg-emerald-100 text-emerald-600 w-16 h-16 rounded-full mx-auto flex items-center justify-center"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-8 w-8"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              ><path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
              /></svg
            >
          </div>
          <h3 class="mt-6 text-xl font-semibold text-gray-900">
            Smart Budgets
          </h3>
          <p class="mt-2 text-gray-600">
            Create budgets that make sense for you and get alerts before you
            overspend.
          </p>
        </div>
        <!-- Feature 3 -->
        <div class="text-center p-8 bg-gray-50 rounded-xl">
          <div
            class="bg-emerald-100 text-emerald-600 w-16 h-16 rounded-full mx-auto flex items-center justify-center"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-8 w-8"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
              ><path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"
              /></svg
            >
          </div>
          <h3 class="mt-6 text-xl font-semibold text-gray-900">
            Achieve Your Goals
          </h3>
          <p class="mt-2 text-gray-600">
            Set savings goals for a vacation or a down payment and track your
            progress.
          </p>
        </div>
      </div>
      <div class="mt-8 flex justify-center space-x-4">
        <a
          href="/register"
          class="bg-emerald-600 text-white px-8 py-3 rounded-lg font-semibold text-lg hover:bg-emerald-700 transition-transform hover:scale-105"
        >
          Get Started for Free
        </a>
      </div>
    </div>
  </section>

  
</main>

<!-- Footer -->
<footer class="bg-white">
  <div class="container mx-auto px-6 py-8">
    <div class="flex justify-between items-center">
      <p class="text-gray-500">&copy; 2025 Basil. All rights reserved.</p>
      <div class="flex space-x-6">
        <a href="/privacy" class="text-gray-500 hover:text-emerald-600"
          >Privacy</a
        >
        <a href="/terms" class="text-gray-500 hover:text-emerald-600">Terms</a>
        <a href="/contact" class="text-gray-500 hover:text-emerald-600"
          >Contact</a
        >
      </div>
    </div>
  </div>
</footer>
