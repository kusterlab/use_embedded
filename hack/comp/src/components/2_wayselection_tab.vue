<template>
<v-card>
  <v-tabs>
    <v-tab
href="#tab-1"
>Item One</v-tab>
    <v-tab
href="#tab-2"
>Item Two</v-tab>
  </v-tabs>
<v-tabs-items v-model="tab">
      <v-tab-item
        v-for="i in 2"
        :key="i"
        :value="'tab-' + i"
>
<v-card flat>
<Table/>
</v-card >
</v-tab-item>
    </v-tabs-items>
</v-card>
</template>

<script>
import Table from './Table.vue'


export default {
	components: {
		Table
	},
  data () {
      return {
        tab: null,
}}
}
</script>
