<template>
  <!-- 导航 -->
  <PcNav class="z-10" />
  <div class="h-full grid grid-cols-24 p-5 overflow-y-scroll relative z-0">
    <LeftBox class="fixed hidden md:block" />
    <main
      class="mx-auto col-start-1 col-span-24 md:col-start-7 md:col-end-25 xl:col-start-6 xl:col-end-20 grid grid-cols-3"
    >
      <NuxtPage class="col-span-3 sm:col-span-2" />
      <AdditionalPanel class="hidden sm:block col-span-1 h-full bg-slate-300" />
    </main>
    <RightBox class="fixed right-5 hidden xl:block" />
  </div>
</template>
