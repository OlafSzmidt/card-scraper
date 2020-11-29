# Card-Scraper

An extremely (and frankly experimental) simple go-lang based webscraper project that checks stock levels of products on Overclockers; alerting to console and SMS if one is present.

# Usage
## Configuration

Configure config.yml with all the target products on Overclockers to track.

For example:
```yaml
sms:
  enabled: true
targets:
  - url: https://www.overclockers.co.uk/pny-nvidia-tesla-volta-v100-16gb-pcie-gpu-accelerator-card-5120-cuda-cores-gx-06a-pn.html
    limit: 5
    selector: input#basketButton
  - url: https://www.overclockers.co.uk/pny-nvidia-tesla-volta-v100-16gb-pcie-gpu-accelerator-card-5120-cuda-cores-gx-06a-pn.html
    limit: 5
    selector: input#basketButton
```

Note; rate limiting is done in seconds. If the product is available every rate limit, and text messages are enabled; a new SMS message will be triggered. This might drain your Twilio balance; adjust the code accordingly.

## Twilio 
If you enable Twilio in the configuration above, set the following environment variables with your API key:

```bash
export ACCOUNTSID=
export AUTHTOKEN=
```

