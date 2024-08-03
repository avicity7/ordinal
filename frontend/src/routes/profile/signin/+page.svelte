<div class="w-screen flex flex-col items-center text-zinc-800">
  <div class="max-w-xl mt-12">
  <h1 class="font-bold text-3xl mb-6 w-full">Sign in</h1>
  <form class="flex flex-col mt-6" method="POST" action="?/login">
    <div class="font-medium mb-4">Email</div>
    <input
      data-testid="sign-in-email-input"
      type="text"
      autocomplete="off"
      name="email"
      class="px-3 border-2 border-zinc-200 rounded-md focus:border-yellow-500 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800"
    />
    <div class="font-medium mb-4 mt-8">Password</div>
    <input
      data-testid="sign-in-password-input"
      type="password"
      autocomplete="off"
      name="password"
      class="px-3 border-2 border-zinc-200 rounded-md focus:border-yellow-500 outline-none select-none h-[2rem] mr-2 dark:text-zinc-800"
    />
    <button class="py-1 bg-primary hover:bg-yellow-700 font-medium text-white rounded-md mt-6 transition-colors" type="submit">Sign in</button>
    <div class="text-zinc-400 font-medium text-sm mx-6 md:mx-24 mt-8 text-center">
      Don't have an account yet?
      <a href="/profile/signup" class="text-blue-600 hover:text-blue-400">
        Sign up
      </a>
    </div>
  	</form>
  </div>
</div>
