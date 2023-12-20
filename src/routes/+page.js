export const ssr = false

export async function load() {
  let resp = await fetch('http://localhost:8999/',
    {
      headers: {
        'Accept': 'application/json'
      }
    })
  let data = await resp.json()
  return { props: { data } }
}
