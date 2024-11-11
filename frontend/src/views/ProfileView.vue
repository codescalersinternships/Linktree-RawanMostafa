<template>
    <NavBar />
    <div class="hero">
        <div class="profile-section">
            <h2 class="name">{{ userInfo.first_name }} {{ userInfo.second_name }}</h2>
            <small class="username">@{{ props.username }}</small>
            <div class="bio">
                <template v-if="!isEditing">
                    <p>{{ userInfo.bio }}</p>
                    <a href="#" @click="toggleEdit">
                        <img class="edit-icon" src="/assets/icons/edit.png" alt="edit icon" />
                    </a>
                </template>
                <template v-else>
                    <textarea v-model="bioInput" class="bio-input"></textarea>
                    <button @click="saveBio" class="save-btn">Save</button>
                    <button @click="toggleEdit" class="cancel-btn">Cancel</button>
                </template>
            </div>
        </div>

        <h3 class="link-title">{{ userInfo.first_name }}'s links:</h3>
        <div v-if="links.length > 0">
            <div v-for="(link, index) in links" :key="index">
                <Link :url="link.url" :platform="link.platform" :click_count="link.click_count" :link_id="link.link_id"  @update-link="updateLink(index, $event)" @delete-link="deleteLink(index, $event)"/>
            </div>
        </div>
        <button class="submit-btn">Add Link</button>
    </div>
</template>

<script setup>
import NavBar from '../components/NavBar.vue';
import Link from '../components/Link.vue';
import { ref, onMounted, reactive } from 'vue';
import axios from 'axios';

const props = defineProps({
    username: {
        type: String,
        required: true
    }
});

const links = ref([]);
const userInfo = reactive({
    id: '',
    bio: '',
    first_name: '',
    second_name: ''
});
const isEditing = ref(false);
const bioInput = ref('');

let token;
onMounted(() => {
    token = localStorage.getItem('authToken');

    axios.get(`http://localhost:8083/api/v1/link/${props.username}`, {
        headers: {
            Authorization: `Bearer ${token}`
        }
    })
    .then(response => {
        links.value = response.data;
    })
    .catch(err => {
        console.error("Error fetching user links:", err);
    });

    axios.get(`http://localhost:8083/user/${props.username}`, {
        headers: {
            Authorization: `Bearer ${token}`
        }
    })
    .then(response => {
        Object.assign(userInfo, response.data);
    })
    .catch(err => {
        console.error("Error fetching user profile:", err);
    });
});

function toggleEdit() {
    isEditing.value = !isEditing.value;
    bioInput.value = userInfo.bio; 
}

function saveBio() {
    axios.put(`http://localhost:8083/user/update-bio/${props.username}`, 
    { "bio": bioInput.value }, 
    {
        headers: {
            Authorization: `Bearer ${token}`
        }
    })
    .then(response => {
        userInfo.bio = bioInput.value; 
        isEditing.value = false;         
    })
    .catch(err => {
        console.error("Error updating bio:", err);
    });
}

function updateLink(index, newLink) { 
  links.value[index] = { ...links.value[index], ...newLink };

  axios.put(`http://localhost:8083/api/v1/link/${newLink.link_id}`, 
  {
    url: newLink.url,
  },
  {
    headers: {
      Authorization: `Bearer ${token}`
    }
  })
  .then(response => {
    console.log("Link updated successfully:", response.data);
  })
  .catch(err => {
    console.error("Error updating link:", err);
  });
}

function deleteLink(index, newLink) { 

  axios.delete(`http://localhost:8083/api/v1/link/${newLink.link_id}`,
  {
    headers: {
      Authorization: `Bearer ${token}`
    }
  })
  .then(response => {
    console.log("Link delete successfully:", response.data);
    links.value.splice(index, 1);

  })
  .catch(err => {
    console.error("Error deleting link:", err);
  });
}
</script>

<style scoped>

.hero{
    position: relative;
    min-height: 82vh;
    font-family: 'Courier New', Courier, monospace;
}
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

.profile-section {
    background-color: rgba(175, 149, 127, 0.454);
    width: 100vw;
    height: fit-content;
    padding-top: 30px;
    padding-left: 30px;
}

.submit-btn:hover {
    background-color: rgb(35, 64, 74);
}

.edit-icon {
    width: 25px;
    margin-left: 10px;
    cursor: pointer;
}

.bio {
    display: flex;
    align-items: center;
}

.bio-input {
    width: 100%;
    padding: 10px;
    margin-top: 10px;
    font-size: 16px;
    border: 1px solid rgb(51, 93, 107);
    border-radius: 5px;
    resize: none;
}

.save-btn, .cancel-btn {
    padding: 5px 10px;
    font-size: 14px;
    margin: 10px 5px;
    background-color: rgb(51, 93, 107);
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}
.username{
    margin-left: 10px;
    color: #504f4f;
}

.save-btn:hover {
    background-color: rgb(35, 64, 74);
}

.cancel-btn {
    background-color: #ccc;
}

.cancel-btn:hover {
    background-color: #999;
}

.link-title {
    margin-left: 20px;
}

.name{
    margin-top: -15px;
    color:  rgb(26, 47, 54);
}
</style>
