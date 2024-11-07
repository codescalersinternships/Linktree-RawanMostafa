<script >
import Register from './components/Register.vue'
import Login from './components/Login.vue'
import Link from './components/Link.vue'
import AddLink from './components/AddLink.vue'

export default {
  components:{
    Register,
    Login,
    Link,
    AddLink,
  },
}
</script>

<template>
  <Register />
  <Login/> 
  <Link />
   <AddLink />
</template>
