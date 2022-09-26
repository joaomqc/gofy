import axios from "axios";

export async function performSearch(
  searchTerm: string,
  type: string
): Promise<{ error: string | null; data: any | null }> {
  try {
    const response = await axios.get(
      `${process.env.VUE_APP_API_URL}/search/${type}/${searchTerm}`
    );

    if (response.status !== 200) {
      return { error: "Request to server failed", data: null };
    }

    return { error: null, data: response.data };
  } catch (e) {
    return { error: "Server could not be reached", data: null };
  }
}

export async function addArtist(
  id: string,
  monitor: boolean
): Promise<string | undefined> {
  try {
    const response = await axios.post(
      `${process.env.VUE_APP_API_URL}/artist/`,
      {
        spotify_id: id,
        monitor,
      }
    );

    if (response.status !== 200) {
      return "Request to server failed";
    }
  } catch (e) {
    return "Server could not be reached";
  }
}

export async function requestArtists(): Promise<{
  error: string | null;
  data: any | null;
}> {
  try {
    const response = await axios.get(`${process.env.VUE_APP_API_URL}/artist/`);

    if (response.status !== 200) {
      return { error: "Request to server failed", data: null };
    }

    return { error: null, data: response.data };
  } catch (e) {
    return { error: "Server could not be reached", data: null };
  }
}
