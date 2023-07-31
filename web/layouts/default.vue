<template>
  <nav
    class="border-b h-16 content-center grid sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-4"
  >
    <div class="md:col-span-1">左</div>
    <div class="hidden lg:block col-span-2">中</div>
    <div class="md:col-span-1 flex items-center">
      <Icon name="icon-xiaoxi" size="text-2xl"></Icon>
      <Icon name="icon-shezhi" size="text-2xl"></Icon>
      <Avatar />
    </div>
  </nav>
  <div>
    <slot />
  </div>
</template>
