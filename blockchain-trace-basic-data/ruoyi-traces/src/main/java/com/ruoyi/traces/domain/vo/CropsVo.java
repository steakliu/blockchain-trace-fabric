package com.ruoyi.traces.domain.vo;

import java.util.Date;

public class CropsVo {
    private Long id;

    private String cropsId;

    private String farmerNickName;

    private String farmerTel;

    private String farmerDept;

    private String cropsName;

    private Date time;

    private Integer machingStatus;

    private Integer outFactoryStatus;

    private Integer productWriteStatus;

    private Integer retailerReceiveStatus;

    public Date getTime() {
        return time;
    }

    public void setTime(Date time) {
        this.time = time;
    }

    public Integer getRetailerReceiveStatus() {
        return retailerReceiveStatus;
    }

    public void setRetailerReceiveStatus(Integer retailerReceiveStatus) {
        this.retailerReceiveStatus = retailerReceiveStatus;
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public Integer getProductWriteStatus() {
        return productWriteStatus;
    }

    public void setProductWriteStatus(Integer productWriteStatus) {
        this.productWriteStatus = productWriteStatus;
    }

    public Integer getMachingStatus() {
        return machingStatus;
    }

    public void setMachingStatus(Integer machingStatus) {
        this.machingStatus = machingStatus;
    }

    public Integer getOutFactoryStatus() {
        return outFactoryStatus;
    }

    public void setOutFactoryStatus(Integer outFactoryStatus) {
        this.outFactoryStatus = outFactoryStatus;
    }

    public String getCropsId() {
        return cropsId;
    }

    public void setCropsId(String cropsId) {
        this.cropsId = cropsId;
    }

    public String getFarmerNickName() {
        return farmerNickName;
    }

    public void setFarmerNickName(String farmerNickName) {
        this.farmerNickName = farmerNickName;
    }

    public String getFarmerTel() {
        return farmerTel;
    }

    public void setFarmerTel(String farmerTel) {
        this.farmerTel = farmerTel;
    }

    public String getFarmerDept() {
        return farmerDept;
    }

    public void setFarmerDept(String farmerDept) {
        this.farmerDept = farmerDept;
    }

    public String getCropsName() {
        return cropsName;
    }

    public void setCropsName(String cropsName) {
        this.cropsName = cropsName;
    }
}
