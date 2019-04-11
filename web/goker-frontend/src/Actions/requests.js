const hostUrl =
  process.env.NODE_ENV === 'production'
    ? process.env.REACT_APP_PROD_API_URL
    : 'http://localhost:8080';

async function getRank(hand) {
  const response = await fetch(`${hostUrl}/?hand=${hand}`, {
    accept: 'application/json',
  });
  const checkedStatus = checkStatus(response);
  const parsedJson = await parseJSON(checkedStatus);
  return parsedJson;
}

function checkStatus(response) {
  if (response.status >= 200 && response.status < 300) {
    return response;
  }
  const error = new Error(`HTTP Error ${response.statusText}`);
  error.status = response.statusText;
  error.response = response;
  console.log(error, 'error'); // eslint-disable-line no-console
  throw error;
}

async function parseJSON(response) {
  return response.json();
}

export {
  getRank
};
