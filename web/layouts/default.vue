<template>
  <nav
    class="shadow-md border-b h-16 content-center px-5 grid grid-cols-2 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-4"
  >
    <div class="md:col-span-1">
      <NuxtLink to="#" class="flex items-center h-full">
        <img
          class="h-8 mr-5"
          src="https://hxsj.in/_nuxt/img/logo.5c704a9.png"
          alt=""
          srcset=""
        />
        KK_BBS
      </NuxtLink>
    </div>
    <div class="hidden lg:flex items-center col-span-2">
      <Search />
      <ul class="w-4/6 h-full flex items-center px-5">
        <li>
          <button class="bg-gray-300 rounded-full w-10 h-10">
            <Icon name="icon-shouye" size="text-2xl mx-1"></Icon>
          </button>
        </li>
      </ul>
    </div>
    <div class="md:col-span-1 flex items-center justify-end">
      <Icon name="icon-xiaoxi" size="text-2xl mx-2"></Icon>
      <Icon name="icon-shezhi" size="text-2xl mx-2"></Icon>
      <Avatar />
    </div>
  </nav>
  <div class="h-full grid grid-cols-12 p-5">
    <LeftBox class="col-span-3 md:col-span-4 lg:col-span-3 hidden sm:block" />
    <main class="col-span-12 md:col-span-8 lg:col-span-6">
      <slot />
    </main>
    <RightBox class="col-span-3 hidden lg:block" />
  </div>
</template>
