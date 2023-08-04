<template>
  <PcNav />
  <div class="h-full grid grid-cols-24 p-5">
    <LeftBox class="hidden sm:block md:col-span-6 lg:col-span-6 xl:col-span-5" />
    <main class="col-span-24 md:col-span-18 lg:col-span-18 xl:col-span-14">
      <NuxtPage />
    </main>
    <RightBox class="col-span-5 hidden xl:flex" />
  </div>
</template>
