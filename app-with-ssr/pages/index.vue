<template>
  <div v-html="blog">
  </div>
</template>

<script>
import axios from 'axios'
import MarkdownIt from 'markdown-it'

export default {
  name: 'App',
  data: () => ({
    blog: ''
  }),
  async asyncData () {
    const { data: response } = await axios.get('http://localhost:3000/my-post.md')
    const md = new MarkdownIt()
    const blog = md.render(response)

    return {
      blog
    }
  }
}
</script>
