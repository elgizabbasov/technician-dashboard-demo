<template>
    <nav>
        <v-toolbar flat>
            <!-- <v-app-bar-nav-icon @click="drawMini = !drawMini"></v-app-bar-nav-icon> -->
            <v-toolbar-title>
                CarbinX Technician
            </v-toolbar-title>
        </v-toolbar>
        <v-navigation-drawer app expand-on-hover :mini-variant="true" height="100%" :clipped="true" color=#dee8ff permanent>
            <v-list nav rounded>
                <router-link to="/dashboard">
                    <v-img src="@/assets/05-CarbinX-Wordmark-TM-DarkBlue-RGB-SM.png"></v-img>
                </router-link>
                <v-divider class="mt-4 mb-4"></v-divider>
                <v-list-item 
                    v-for="item in filteredItems" 
                    :key="item.name" 
                    :to="item.route" 
                    link
                >
                    <v-icon>{{ item.icon }}</v-icon>
                    <v-list-item-title class="ml-4">{{ item.name }}</v-list-item-title>
                </v-list-item>
                <v-list-item v-if="selectDeviceItem" @click="dialog = true">
                    <v-icon>{{ selectDeviceItem.icon }}</v-icon>
                    <v-list-item-title class="ml-4">{{ selectDeviceItem.name }}</v-list-item-title>
                </v-list-item>
            </v-list>
            <v-dialog v-model="dialog">
                <select-device-component/>
                <v-btn
                    color=#4d00fa
                    class="white--text"
                    @click="dialog = false"
                >
                    Done
              </v-btn>
            </v-dialog>
        </v-navigation-drawer>
    </nav> 
</template>

<script>
export default {
  name: "App",
  data: () => ({
    items: [
      {name: 'Dashboard', icon: 'mdi-view-dashboard', route: '/dashboard', type: 'route'},
      {name: 'Checkup Form', icon: 'mdi-form-select', route: '/form', type: 'route'},
      {name: 'Setup', icon: 'mdi-cogs', route: '/setup', type: 'route'},
      {name: 'Select New Device', icon: 'mdi-devices', type: 'dialog'},

    ],
    draw: true,
    drawMini: false,
    dialog: false,
  }),
  computed: {
    selectDeviceItem() {
      return this.items.find(item => item.type === 'dialog');
    }, 
    filteredItems() {
      return this.items.filter(item => item.type === 'route');
    }
  }
};
</script>
  
<!-- <v-navigation-drawer app permanent color=#becbf6>
    <v-list-item>
        <v-list-item-content>
            <v-list-item-title class="text-h5"> CarbinX Technician </v-list-item-title>
        </v-list-item-content>
    </v-list-item>
    <v-divider color="black" thickness="20px"></v-divider>
    <v-list nav rounded>
        <v-list-item v-for="item in items" :key="item.name" router :to="item.route" link>
            <v-list-item-content>
                <v-list-item-title>{{ item.name }}</v-list-item-title>
            </v-list-item-content>
        </v-list-item>
    </v-list>
    <v-container fluid style="margin: 0px; padding: 0px; width: 100%">
        <select-device-component/>
    </v-container>
</v-navigation-drawer> -->