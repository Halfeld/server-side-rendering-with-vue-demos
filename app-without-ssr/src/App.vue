<template>
  <div id="app">
    <div v-html="blog"></div>
  </div>
</template>

<script>
import axios from 'axios'
import MarkdownIt from 'markdown-it'

export default {
  name: 'app',
  data: () => ({
    blog: ''
  }),
  async mounted () {
    const { data: response } = await axios.get('static/my-post.md')
    const md = new MarkdownIt()
    this.blog = md.render(response)
  }
}
</script>
