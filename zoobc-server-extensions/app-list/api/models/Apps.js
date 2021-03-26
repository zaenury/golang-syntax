const mongoose = require('mongoose')

const schema = new mongoose.Schema(
  {
    Name: { type: String, required: '{PATH} must be completed!' },
    Api: { type: String, required: '{PATH} must be completed!' },
    Active: { type: Boolean, required: '{PATH} must be completed!' },
    Order: { type: Number },
    Desciption: { type: String },
  },
  {
    timestamps: true,
    toJSON: { virtuals: true },
  }
)

module.exports = mongoose.model('Apps', schema)
