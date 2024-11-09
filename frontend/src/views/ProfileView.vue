<template>
    <NavBar />
    <div class="hero">
        <h1>User Profile: {{ username }}</h1>
        <div v-if="links.length > 0">
            <div v-for="(link, index) in links" :key="index">
                <Link :url="link.url" :platform="link.platform" :click_count="link.click_count" />
            </div>
        </div>
        <div v-else>
            <p>Loading...</p>
        </div>
        <button class="submit-btn">Add Link</button>
    </div>
</template>

<script setup>
import NavBar from '../components/NavBar.vue'
import Link from '../components/Link.vue'
import { ref, onMounted } from 'vue'
import axios from 'axios'

const props = defineProps({
    username: {
        type: String,
        required: true
    }
})

const links = ref([])

onMounted(() => {
    const token = localStorage.getItem('authToken');

    axios.get(`http://localhost:8083/api/v1/link/${props.username}`, {
        headers: {
            Authorization: `Bearer ${token}`
        }
    })
        .then(response => {
            links.value = response.data
            console.log(response.data)
        })
        .catch(err => {
            console.error("Error fetching user profile:", err)
        })
})
</script>

<style scoped>
.submit-btn {
    padding: 10px 20px;
    font-size: 16px;
    margin: 20px;
    background-color: rgb(51, 93, 107);
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}

.submit-btn:hover {
    background-color: rgb(35, 64, 74);
    
}
.hero {
    padding-top: 100px; 
}
</style>