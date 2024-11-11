<template>
    <div class="link">
      <div class="hero-link">
        <img class="platform-logo" :src="`/assets/icons/${platform}.png`" :alt="`${platform} logo`" />
        <template v-if="!isEditing"> 
          <a :href="url" target="_blank">{{ url }}</a>
          <div class="dropdown">
            <div class="dropdown-btn">...</div>
            <div class="dropdown-content">
              <button class="dropdown-item" @click="toggleEdit">Edit</button> 
              <button class="dropdown-item">Delete</button>
            </div>
          </div>
        </template>
        <template v-else> 
          <input v-model="editableUrl" placeholder="Edit URL" /> 
          <button @click="saveEdit">Save</button>
          <button @click="toggleEdit">Cancel</button> 
        </template>
      </div>
      <div class="click-count">
        <small>Click count: {{ click_count }}</small>
      </div>
    </div>
  </template>

<script setup>
import {ref} from "vue";
const props=defineProps({
  url: {
    type: String,
    required: true
  },
  platform: {
    type: String,
    required: true
  },
  click_count: {
    type: Number,
    required: true
  },
  link_id: {
    type: String,
    required: true
  }
})
const isEditing = ref(false); 
const editableUrl = ref(props.url); 

function toggleEdit() {
  isEditing.value = !isEditing.value;
  editableUrl.value = url;
}
const emit = defineEmits(['update-link']);

function saveEdit() { 
  isEditing.value = false;
  emit('update-link', { url: editableUrl.value ,link_id:props.link_id});
  console.log(props.link_id)
}
</script>

<style>
html,
body {
    height: 100%;
    margin: 0;
    font-family: Arial, sans-serif;
}

body {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    background-color: rgba(209, 243, 240, 0.634);
    padding: 10px;
}
.platform-logo{
    width: 30px;
    margin-top: 10px;
}

.link {
    width: 100%;
    max-width: 400px;
    padding: 15px 20px ;
    border: 3px solid rgb(51, 93, 107);
    border-radius: 20px;
    display: flex;
    flex-direction: column; 
    align-items: flex-start;
    background-color: rgba(147, 174, 172, 0.634);
    margin: 20px;
}

.hero-link {
    width: 100%;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
}

.link a {
    color: rgb(36, 73, 70);
    font-size: 18px;
    margin-top: 10px;
    text-decoration: none;
}

.dropdown {
    position: relative;
    display: inline-block;
}

.dropdown-btn {
    width: 30px;
    height: 10px; 
    font-size: 20px;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 5px;
    border: 1px solid rgb(51, 93, 107);
    background-color: rgba(209, 243, 240, 0.7);
    color: rgb(36, 73, 70);
    cursor: pointer;
    padding-bottom: 4px;
    font-weight: bolder;
}

.dropdown-content {
    display: none;
    position: absolute;
    right: 0;
    background-color: white;
    border: 1px solid rgb(51, 93, 107);
    border-radius: 5px;
    box-shadow: 0px 8px 16px 0px rgba(0, 0, 0, 0.2);
    z-index: 1;
    min-width: 100px;
}

.dropdown-content .dropdown-item {
    background-color: white;
    color: rgb(36, 73, 70);
    padding: 10px;
    text-align: left;
    border: none;
    width: 100%;
    cursor: pointer;
}

.dropdown-content .dropdown-item:hover {
    background-color: rgba(209, 243, 240, 0.5);
}

.dropdown:hover .dropdown-content {
    display: block;
}

.click-count {
    margin-top: 10px;
    color: rgb(36, 73, 70);
    font-size: 17px;
    align-self: flex-start;
    margin-bottom: 5px;
}

.link a:hover{
    color: rgb(15, 40, 38);
}
</style>
