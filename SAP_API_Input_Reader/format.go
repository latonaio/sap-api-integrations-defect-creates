package sap_api_input_reader

import (
	"sap-api-integrations-defect-creates/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.Defect
	return &requests.Header{
		DefectInternalID:            data.DefectInternalID,           
		Defect:                      data.Defect,                     
		DefectCategory:              data.DefectCategory,             
		CreationDate:                data.CreationDate,               
		LastChangeDate:              data.LastChangeDate,             
		DefectText:                  data.DefectText,                 
		DefectCodeCatalog:           data.DefectCodeCatalog,          
		DefectCodeGroup:             data.DefectCodeGroup,            
		DefectCode:                  data.DefectCode,                 
		DefectCodeVersion:           data.DefectCodeVersion,          
		DefectObjectCodeCatalog:     data.DefectObjectCodeCatalog,    
		DefectObjectCodeGroup:       data.DefectObjectCodeGroup,      
		DefectObjectCode:            data.DefectObjectCode,           
		DefectiveQuantity:           data.DefectiveQuantity,          
		DefectiveQuantityUnit:       data.DefectiveQuantityUnit,      
		ManufacturingOrder:          data.ManufacturingOrder,         
		OrderInternalID:             data.OrderInternalID,            
		ManufacturingOrderOperation: data.ManufacturingOrderOperation,
		ManufacturingOrderSequence:  data.ManufacturingOrderSequence, 
		CreationTime:                data.CreationTime,               
		LastChangeTime:              data.LastChangeTime,             
		DefectClass:                 data.DefectClass,                
		NumberOfDefects:             data. NumberOfDefects,            
		InspPlanOperationInternalID: data.InspPlanOperationInternalID,
		InspectionCharacteristic:    data.InspectionCharacteristic,   
		InspectionSubsetInternalID:  data.InspectionSubsetInternalID, 
		MaterialSample:              data.MaterialSample,             
		WorkCenterTypeCode:          data.WorkCenterTypeCode,         
		MainWorkCenterInternalID:    data.MainWorkCenterInternalID,   
		MainWorkCenterPlant:         data.MainWorkCenterPlant,        
		MainWorkCenter:              data.MainWorkCenter,             
		Equipment:                   data.Equipment,                  
		FunctionalLocation:          data.FunctionalLocation,         
		IsDeleted:                   data.IsDeleted,                  
		DefectOrigin:                data.DefectOrigin,               
		Material:                    data.Material,                   
		Plant:                       data.Plant,                      
		InspectionLot:               data.InspectionLot,              
		CatalogProfile:              data.CatalogProfile,             
		ChangedDateTime:             data.ChangedDateTime,            
	}
}