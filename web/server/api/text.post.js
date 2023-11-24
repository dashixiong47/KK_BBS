export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  let res = await fetch(body.url)
  let text=await res.text()
  return text
})
