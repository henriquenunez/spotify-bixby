var http = require('http');
var config = require('config')

module.exports.function = function searchArtist (artist_name) {
  var url = config.get('searchEndpoint');
  
  var response = http.getUrl(
    url,
    {
      format: 'json',
      query: {
        q: artist_name,
        type: 'album',
      }
    }
  );

  //var name = response.artists.items[0].uri + ":play"; //So it plays automatically
  var name = response.albums.items[0].uri + ":play"; //So it plays automatically
  
  return name; //Name is a text. good
}
