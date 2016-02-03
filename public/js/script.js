var URL_PATTERN = /^(https?:\/\/)?([a-zA-Z0-9\-]{1,63}\.)?([a-zA-Z0-9\-]{1,253}\.)([a-zA-Z0-9\-]{2,24}\.)?([a-zA-Z0-9\-]{2,24})(([/]{1})(.)*?)?$/;



function isValid(url) {
   return URL_PATTERN.test(url);
}

function testForURL(textBox) {
   var urlBoxText = textBox.value;
   var formDiv = document.getElementById('form-div');
   var valid = isValid(urlBoxText);

   if (valid) {
      formDiv.classList.add("has-success");
      formDiv.classList.remove("has-error");
   } else {
      formDiv.classList.add("has-error");
      formDiv.classList.remove("has-success");
   }

   return valid;
}

function sendRequest() {
  var form = document.getElementById('url-shortener-form');
  var textBox = document.getElementById('url-box');

  if (!isValid(textBox)) return false;

  $.ajax({
    url: '/api/shorten',
    data: {
      longUrl: textBox.value
    },
    type: 'POST',
  }).fail(function() {alert("FAILURE");}).always(function() {
    alert("Done");
  });

  return true;
}

function test() {
   alert("test!");
}
