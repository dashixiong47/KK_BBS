<template>
  <PcNav />
  <div class="h-full grid grid-cols-12 p-5">
    <LeftBox class="col-span-3 md:col-span-4 lg:col-span-3 hidden sm:block" />
    <main class="col-span-12 md:col-span-8 lg:col-span-6">
      <NuxtPage />
    </main>
    <RightBox class="col-span-3 hidden lg:block" />
  </div>
</template>
