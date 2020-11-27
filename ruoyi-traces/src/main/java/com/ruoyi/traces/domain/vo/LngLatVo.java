package com.ruoyi.traces.domain.vo;

public class LngLatVo {
    private Long id;
    private String cropsId;
    private Long lng;
    private Long lat;

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public String getCropsId() {
        return cropsId;
    }

    public void setCropsId(String cropsId) {
        this.cropsId = cropsId;
    }

    public Long getLng() {
        return lng;
    }

    public void setLng(Long lng) {
        this.lng = lng;
    }

    public Long getLat() {
        return lat;
    }

    public void setLat(Long lat) {
        this.lat = lat;
    }

    @Override
    public String toString() {
        return "LngLatVo{" +
                "cropsId='" + cropsId + '\'' +
                ", lng=" + lng +
                ", lat=" + lat +
                '}';
    }
}
