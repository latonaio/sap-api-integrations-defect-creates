package sap_api_caller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sap-api-integrations-defect-creates/SAP_API_Caller/requests"
	sap_api_output_formatter "sap-api-integrations-defect-creates/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	sap_api_post_header_setup "github.com/latonaio/sap-api-post-header-setup"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	postClient      *sap_api_post_header_setup.SAPPostClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, postClient *sap_api_post_header_setup.SAPPostClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		postClient:      postClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncPostDefect(
	header *requests.Header,
	defectText *requests.Header,	
	accepter []string) {
		wg := &sync.WaitGroup{}
		wg.Add(len(accepter))
		for _, fn := range accepter {
			switch fn {
			case "Header":
				func() {
					c.Header(header)
					wg.Done()
				}()
			case "DefectText":
				func() {
					c.DefectText(defectText)
					wg.Done()
				}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(header *requests.Header) {
	outputDataHeader, err := c.callDefectSrvAPIRequirementHeader("A_Defect", header)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataHeader)
}

func (c *SAPAPICaller) callDefectSrvAPIRequirementHeader(api string, header *requests.Header) (*sap_api_output_formatter.Header, error) {
	body, err := json.Marshal(header)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.postClient.POST(url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}

	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) DefectText(defectText *requests.Header) {
	outputDataDefectText, err := c.callDefectSrvAPIRequirementDefectText("A_ProductPlantDefectTextMgmt", defectText)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(outputDataDefectText)
}

func (c *SAPAPICaller) callDefectSrvAPIRequirementDefectText(api string, defectText *requests.Header) (*sap_api_output_formatter.Header, error) {
	body, err := json.Marshal(defectText)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	url := strings.Join([]string{c.baseURL, "API_PRODUCT_SRV", api}, "/")
	params := c.addQuerySAPClient(map[string]string{})
	resp, err := c.postClient.POST(url, params, string(body))
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, xerrors.Errorf("bad response:%s", string(byteArray))
	}
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) addQuerySAPClient(params map[string]string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["sap-client"] = c.sapClientNumber
	return params
}
