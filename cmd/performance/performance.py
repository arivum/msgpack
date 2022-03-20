#%%
import json
import pandas
import altair
import altair_saver
import os

def read_and_transform(jsonFile):
    data = []
    perfJson = json.load(open(jsonFile))
    for test in perfJson:
        #print(test["Stats"])
        data.extend([ { "name": test["Name"], "value": x["MBPerSec"] } for x in test["Stats"][:-1] ])
    data = pandas.DataFrame(data)
    print(data)
    return data


# %%
altair.data_transformers.disable_max_rows()

def draw(data, title):
    return altair.Chart(data).mark_boxplot(extent="min-max", size=30).encode(
        x=altair.X('name:N', axis=None),
        y=altair.Y('value:Q', title="Throughput [NB/s]", scale=altair.Scale(domain=[0, 150])),
        color=altair.Color('name:N', title="Implementation", scale=altair.Scale(scheme='darkmulti'))
    ).properties(
        width=400, height=600
    ).configure_axis(
        labelFontSize=13,
        titleFontSize=16
    ).configure_legend(
        labelFontSize=13,
        titleFontSize=16
    )


# %%
chart = draw(read_and_transform("../performance.json"), "Performance")
chart
altair_saver.save(chart, "performance.png")
os.unlink("geckodriver.log")

# %%
