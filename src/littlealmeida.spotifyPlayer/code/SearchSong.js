var http = require('http');
var config = require('config')

module.exports.function = function searchArtist (song_name, artist_name) {
  var url = config.get('searchEndpoint')
  
  var response = http.getUrl(
    url,
    {
      format: 'json',
      query: {
        song: song_name,// + ' ' + artist_name,
        //type: 'track',
      }
    }
  );

  var uri = response.uri + ":play"; //So it plays automatically
  
  return uri;
}
