<template>
    <div class="add-link">
        <h2>Add a New Link</h2>
        <div class="platforms">
            <img v-for="platform in platforms" :key="platform.name" :src="platform.icon" :alt="platform.name"
                :class="{ selected: selectedPlatform === platform.name }" @click="selectPlatform(platform.name)"
                class="platform-icon" />
        </div>
        <input type="url" placeholder="Enter URL here" v-model="url" class="url-input" />
        <button @click="addLink" :disabled="!url || !selectedPlatform">Add Link</button>
    </div>
</template>

<script>
import igIcon from '../assets/icons/instagram.png';
import fbIcon from '../assets/icons/facebook.png';
import xIcon from '../assets/icons/twitter.png';
import linkedinIcon from '../assets/icons/twitter.png'
import gitIcon from '../assets/icons/github.png'


export default {
    data() {
        return {
            platforms: [
                { name: "Instagram", icon: igIcon },
                { name: "Twitter", icon: xIcon },
                { name: "Facebook", icon: fbIcon },
                { name: "Linkedin", icon: linkedinIcon },
                { name: "Github", icon: gitIcon },
            ],
            selectedPlatform: null,
            url: ""
        };
    },
    methods: {
        selectPlatform(platformName) {
            this.selectedPlatform = platformName;
        },
        addLink() {
            this.selectedPlatform = null;
            this.url = "";
        }
    }
};
</script>

<style scoped>
html,
body {
    height: 100%;
    margin: 0;
    font-family: Arial, sans-serif;
}

.add-link {
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 100%;
    max-width: 400px;
    padding: 20px;
    border: 3px solid rgb(51, 93, 107);
    border-radius: 20px;
    background-color: rgba(147, 174, 172, 0.8);
    box-shadow: 0px 4px 8px rgba(0, 0, 0, 0.2);
}

h2 {
    color: rgb(36, 73, 70);
    font-size: 20px;
    margin-bottom: 10px;
    text-align: center;
}

.platforms {
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    gap: 10px;
    margin-bottom: 15px;
}

.platform-icon {
    height: 40px;
    cursor: pointer;
    border: 4px solid transparent;
    border-radius: 10px;
    transition: transform 0.2s, border-color 0.2s;
}

.platform-icon.selected {
    border-color: rgb(51, 93, 107);
    border-width: 4px;
    transform: scale(1.1);
}

.url-input {
    width: 100%;
    padding: 10px;
    font-size: 16px;
    border: 2px solid rgb(51, 93, 107);
    border-radius: 5px;
    margin-bottom: 10px;
    outline: none;
}

button {
    padding: 10px 20px;
    font-size: 16px;
    color: white;
    background-color: rgb(51, 93, 107);
    border: none;
    border-radius: 5px;
    cursor: pointer;
    transition: background-color 0.2s;
}

button:disabled {
    background-color: gray;
    cursor: not-allowed;
}

button:not(:disabled):hover {
    background-color: rgb(36, 73, 70);
}



@media (max-width: 500px) {
    .platform-icon {
        width: 40px;
        height: 40px;
    }

    h2 {
        font-size: 18px;
    }

    .url-input,
    button {
        font-size: 14px;
    }
}
</style>
