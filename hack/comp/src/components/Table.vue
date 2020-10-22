<template>
<v-card>
    <v-card-title>
      Nutrition
      <v-spacer></v-spacer>
      <v-text-field
        v-model="search"
        append-icon="mdi-magnify"
        label="Search"
        single-line
        hide-details
      ></v-text-field>
    </v-card-title>
  <v-data-table
    v-model="selected"
    :headers="headers"
    :items="api"
    :single-select="singleSelect"
    item-key="usi"
    show-select
    class="elevation-1"
	:search="search"
  >
  </v-data-table>
</v-card>
</template>

<script>
import axios from 'axios';

  export default {
props: ['tableId', 'peptideSequence'],
    data: () => ({
        singleSelect: true,
        selected: [],
	search: '',
        headers: [
          {
            text: 'Usi',
            align: 'start',
            sortable: true,
            value: 'usi',
          },
          { text: 'PeptideSequence', value: 'peptideSequence' },
          { text: 'Charge', value: 'charge' },
          { text: 'precursorMz', value: 'precursorMz' },
          { text: 'valid', value: 'valid' },
          { text: 'Links', value: 'links.self.href' }
        ],
      api: [],
    }),
mounted () {
	var that = this;
console.log(that.peptideSequence);
    axios
      .get('https://www.ebi.ac.uk/pride/ws/archive/v2/spectra?peptideSequence=' + that.peptideSequence + '&pageSize=1000')
      .then(function(response){
		console.log(response.data._embedded.spectraevidences);
		that.api = response.data._embedded.spectraevidences;
		}
);
console.log(that.tableId);
      
  }
  }
</script>
